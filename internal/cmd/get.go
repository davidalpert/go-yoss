package cmd

import (
	"fmt"
	"github.com/davidalpert/go-printers/v1"
	"github.com/davidalpert/go-yoss/internal/cfgset"
	"github.com/davidalpert/go-yoss/internal/provider"
	"github.com/davidalpert/go-yoss/internal/provider/paramstore"
	"github.com/spf13/cobra"
	"sort"
	"strings"
)

type GetOptions struct {
	*printers.PrinterOptions
	provider.Options
	SupportedProviders map[string]provider.NewProviderFn
	Key                string
	Recursive          bool
	//DecryptResult bool
}

func NewGetOptions(ioStreams printers.IOStreams) *GetOptions {
	return &GetOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(ioStreams).WithDefaultOutput("text"),
		SupportedProviders: map[string]provider.NewProviderFn{
			paramstore.ProviderKey: paramstore.NewProvider,
		},
	}
}

func NewCmdGet(ioStreams printers.IOStreams) *cobra.Command {
	o := NewGetOptions(ioStreams)
	var cmd = &cobra.Command{
		Use:     "get <provider> <path>",
		Short:   "get a config value",
		Aliases: []string{"g"},
		Args:    cobra.ExactArgs(2),
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

	cmd.Flags().BoolVarP(&o.Debug, "debug", "d", false, "enable debug output")
	cmd.Flags().BoolVarP(&o.Recursive, "recursive", "r", false, "recursively get all values under that path")

	return cmd
}

func (o *GetOptions) supportedProviderKeys() []string {
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
func (o *GetOptions) Complete(cmd *cobra.Command, args []string) error {
	o.ProviderName = args[0]

	if len(args) > 1 {
		o.Key = args[1]
	} else {
		o.Key = "/"
	}

	return nil
}

// Validate the options
func (o *GetOptions) Validate() error {
	if _, ok := o.SupportedProviders[o.ProviderName]; !ok {
		return fmt.Errorf("unrecognized provider %#v: supported provider are: %#v", o.ProviderName, strings.Join(o.supportedProviderKeys(), ", "))
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *GetOptions) Run() error {
	p, err := o.SupportedProviders[o.ProviderName](&o.Options)
	if err != nil {
		return fmt.Errorf("building provider: %s", err)
	}

	flattened := map[string]string{}
	if o.Recursive {
		if many, err := p.GetValueTree(o.Key); err != nil {
			return err
		} else {
			for k, v := range many {
				flattened[k] = v
			}
		}
	} else {
		if one, err := p.GetValue(o.Key); err != nil {
			return err
		} else {
			flattened[o.Key] = one
		}
	}

	if o.FormatCategory() == "text" {
		return o.WriteOutput(cfgset.FlattenedToString(flattened))
	}
	return o.WriteOutput(flattened)
}
