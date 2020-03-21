package cmd

import (
	"github.com/alimy/kr/internal/tars"
	"github.com/spf13/cobra"
)

var (
	outDir      string
	tarsPath    string
	importPaths []string
)

func init() {
	tarsCmd := &cobra.Command{
		Use:   "tars",
		Short: "generate tars code from *.tars",
		Long:  "generate tars code from *.tars",
		Run:   tarsRun,
	}

	tarsCmd.Flags().StringVarP(&outDir, "outdir", "o", "", "which dir to put generated code")
	tarsCmd.Flags().StringVarP(&tarsPath, "tarsPath", "p", "github.com/TarsCloud/TarsGo/tars", "specify the tars source path")
	tarsCmd.Flags().StringSliceVarP(&importPaths, "import", "I", []string{}, "Specify a specific import path")
	register(tarsCmd)
}

func tarsRun(cmd *cobra.Command, _args []string) {
	for _, filename := range cmd.Flags().Args() {
		gen := tars.NewGenGo(filename, outDir, tarsPath, importPaths)
		gen.Gen()
	}
}
