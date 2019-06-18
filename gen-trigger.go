package flogo-dev-tool

import (
	
	"github.com/spf13/cobra"
)
var genTrigger = &cobra.Command{
	Use:   "gen-trigger",
	Short: "Generate activity scaffold",
	Long:  `This subcommand helps you generate activity-scaffold`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		
		err := os.Mkdir(args[0],os.ModePerm)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating dir: %v\n", err)
			os.Exit(1)
		}

		pwd, err := os.Getwd()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting current dir: %v\n", err)
			os.Exit(1)
		}

	
		err = copyFiles(filepath.Join(COREPATH,"trigger"), filepath.Join(pwd,args[0]))

	},
}

func init() {
	descCmd.AddCommand(genTrigger)
}