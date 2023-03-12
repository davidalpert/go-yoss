package cmd

import (
	"fmt"
	"github.com/davidalpert/go-printers/v1"
	"github.com/davidalpert/go-yoss/internal/app"
	"github.com/davidalpert/go-yoss/internal/cfgset"
	"github.com/davidalpert/go-yoss/internal/provider"
	"github.com/davidalpert/go-yoss/internal/provider/paramstore"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"sort"
	"strings"
)

func NewCmdSync(ioStreams printers.IOStreams) *cobra.Command {
	o := NewSyncProviderOptions(ioStreams)
	var cmd = &cobra.Command{
		Use:     "sync <provider>",
		Aliases: []string{"s"},
		Short:   "synchronize merged configs with the given provider",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.Complete(cmd, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			return o.Run()
		},
	}

	o.PrinterOptions.AddPrinterFlags(cmd.Flags())
	o.Options.AddProviderOptions(cmd.Flags())

	cmd.Flags().BoolVarP(&o.Options.Debug, "debug", "d", false, "enable debug output")
	cmd.Flags().BoolVar(&o.DryRun, "dry-run", false, "dry run (don't update remote)")
	cmd.Flags().StringVarP(&o.SourceFolder, "source-folder", "s", ".", "source folder")
	cmd.Flags().StringVarP(&o.KeyPrefix, "prefix", "p", "/", "key prefix")

	return cmd
}

type SyncProviderOptions struct {
	*printers.PrinterOptions
	cfgset.MergeOptions
	provider.Options
	SupportedProviders map[string]provider.NewProviderFn
	KeyPrefix          string
	DryRun             bool
}

func NewSyncProviderOptions(ioStreams printers.IOStreams) *SyncProviderOptions {
	return &SyncProviderOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(ioStreams).WithDefaultTableWriter(),
		SupportedProviders: map[string]provider.NewProviderFn{
			paramstore.ProviderKey: paramstore.NewProvider,
		},
	}
}

func (o *SyncProviderOptions) supportedProviderKeys() []string {
	result := make([]string, len(o.SupportedProviders))
	i := 0
	for key, _ := range o.SupportedProviders {
		result[i] = key
		i++
	}
	sort.Strings(result)
	return result

}

// Complete the options
func (o *SyncProviderOptions) Complete(cmd *cobra.Command, args []string) error {
	o.MergeOptions.Debug = o.Options.Debug
	o.ProviderName = args[0]
	return nil
}

// Validate the options
func (o *SyncProviderOptions) Validate() error {
	if _, ok := o.SupportedProviders[o.ProviderName]; !ok {
		return fmt.Errorf("unrecognized provider %#v: supported provider are: %#v", o.ProviderName, strings.Join(o.supportedProviderKeys(), ", "))
	}

	return o.PrinterOptions.Validate()
}

type SyncKeyResult struct {
	Key         string  `json:"key,omitempty"`
	NewValue    string  `json:"new_value"`
	OldValue    *string `json:"old_value,omitempty"`
	ActionTaken string  `json:"action_taken"`
}

// Run the command
func (o *SyncProviderOptions) Run() error {
	p, err := o.SupportedProviders[o.ProviderName](&o.Options)
	if err != nil {
		return fmt.Errorf("building provider: %s", err)
	}

	mergeResults, err := cfgset.Merge(o.MergeOptions)
	if err != nil {
		return err
	}

	flattened := make(map[string]string)
	for _, r := range mergeResults {
		for k, v := range r.FlattenToMap() {
			flattened[o.KeyPrefix+k] = v
		}
	}

	remoteValuesByKey, err := p.GetValueTree(o.KeyPrefix)
	if err != nil {
		return err
	}

	syncResults := make(map[string]SyncKeyResult)

	for k, v := range flattened {
		keyResult := SyncKeyResult{
			//Key:      k,
			NewValue: v,
		}
		if ov, found := app.LookupByKeyEqualFold(remoteValuesByKey, k); found {
			keyResult.OldValue = &ov
		}
		if keyResult.OldValue != nil && *keyResult.OldValue == v {
			keyResult.ActionTaken = "none: values match"
		} else if o.DryRun {
			keyResult.ActionTaken = "none: (needs update)"
		} else {
			err = p.SetValue(k, v)
			if err != nil {
				keyResult.ActionTaken = err.Error()
			} else {
				keyResult.ActionTaken = "updated"
			}
		}

		syncResults[k] = keyResult
	}

	return o.WithTableWriter("sync results", func(t *tablewriter.Table) {
		t.SetHeader([]string{"Key", "Old Value", "New Value", "Action Taken"})
		keys := make([]string, 0)
		for k, _ := range syncResults {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			r := syncResults[k]
			oldValue := "<none>"
			if r.OldValue != nil {
				oldValue = *r.OldValue
			}
			t.Append([]string{k, oldValue, r.NewValue, r.ActionTaken})
		}

	}).WriteOutput(syncResults)
	//
	//if err = app.Fs.MkdirAll(o.OutFolder, os.ModePerm); err != nil {
	//	return fmt.Errorf("making %#v: %#v", o.OutFolder, err)
	//}
	//
	//for _, appResult := range result {
	//	appOutDir := path.Join(o.OutFolder, appResult.AppDir)
	//	if err = app.Fs.MkdirAll(appOutDir, os.ModePerm); err != nil {
	//		return fmt.Errorf("making %#v: %#v", appOutDir, err)
	//	}
	//
	//	for slug, mergeResult := range appResult.MergeBySlug {
	//		outFile := path.Join(appOutDir, fmt.Sprintf("%s.%s", slug, o.OutFormat))
	//		b, err := yaml.Marshal(mergeResult)
	//		if err != nil {
	//			return fmt.Errorf("marshalling %#v to %#v: %#v", mergeResult, outFile, err)
	//		}
	//
	//		if err = afero.WriteFile(app.Fs, outFile, b, os.ModePerm); err != nil {
	//			return fmt.Errorf("writing %#v: %#v", outFile, err)
	//		}
	//
	//		// TODO: collect errors into an error result rather than failing out on the first one and write to STDERR
	//	}
	//}

	//return o.WithDefaultOutput("json").WriteOutput(result)
	//return nil
}
