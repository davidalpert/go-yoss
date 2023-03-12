package cmd

import (
	"fmt"
	"github.com/davidalpert/go-printers/v1"
	"github.com/davidalpert/go-yoss/internal/app"
	"github.com/davidalpert/go-yoss/internal/cfgset"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"strings"
)

type MergeFolderOptions struct {
	*printers.PrinterOptions
	cfgset.MergeOptions
	OutDir string
}

func NewMergeFolderOptions(ioStreams printers.IOStreams) *MergeFolderOptions {
	return &MergeFolderOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(ioStreams).WithDefaultOutput("yaml"),
	}
}

func NewCmdMergeFolder(ioStreams printers.IOStreams) *cobra.Command {
	o := NewMergeFolderOptions(ioStreams)
	var cmd = &cobra.Command{
		Use:     "folder <source_folder>",
		Short:   "merge two config files together",
		Aliases: []string{"f", "fs"},
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

	cmd.Flags().BoolVarP(&o.Debug, "debug", "d", false, "enable debug output")
	cmd.Flags().StringVar(&o.OutDir, "out-dir", "out", "folder to place output")

	return cmd
}

// Complete the options
func (o *MergeFolderOptions) Complete(cmd *cobra.Command, args []string) error {
	o.SourceFolder = args[0]
	return nil
}

// Validate the options
func (o *MergeFolderOptions) Validate() error {
	if strings.EqualFold(o.SourceFolder, o.OutDir) {
		return fmt.Errorf("cannot merge into the source folder '%s'; please specify a different --out-dir", o.SourceFolder)
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *MergeFolderOptions) Run() error {
	result, err := cfgset.Merge(o.MergeOptions)
	if err != nil {
		return err
	}

	if err = app.Fs.MkdirAll(o.OutDir, os.ModePerm); err != nil {
		return fmt.Errorf("making %#v: %#v", o.OutDir, err)
	}

	outputExt := o.FormatCategory()
	if outputExt == "text" {
		outputExt = "txt"
	}

	for _, appResult := range result {
		appOutDir := path.Join(o.OutDir, appResult.AppDir)
		if err = app.Fs.MkdirAll(appOutDir, os.ModePerm); err != nil {
			return fmt.Errorf("making %#v: %#v", appOutDir, err)
		}

		for slug, mergeResult := range appResult.MergeBySlug {
			outFile := path.Join(appOutDir, fmt.Sprintf("%s.%s", slug, outputExt))
			b, err := yaml.Marshal(mergeResult)
			if err != nil {
				return fmt.Errorf("marshalling %#v to %#v: %#v", mergeResult, outFile, err)
			}

			if err = afero.WriteFile(app.Fs, outFile, b, os.ModePerm); err != nil {
				return fmt.Errorf("writing %#v: %#v", outFile, err)
			}

			// TODO: collect errors into an error result rather than failing out on the first one and write to STDERR
		}
	}

	return nil
}
