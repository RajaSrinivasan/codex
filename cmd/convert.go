package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/RajaSrinivasan/codex/impl/convert"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert code fragments from the input file into graphic file(s)",
	Long: `
	Code fragments are extracted from the specified file and
	a graphic file is generated for each fragment.

	usage:
		convert sourcefile graphicbase linefrom:lines linefrom:lines ...

	output files are named in the format graphicbase.<fragment no>.png
	
	`,
	Args: cobra.MinimumNArgs(3),
	Run:  Convert,
}

func init() {
	rootCmd.AddCommand(convertCmd)
}

func Convert(cmd *cobra.Command, args []string) {
	convert.Convert(args[0], args[1], args[2:])
}
