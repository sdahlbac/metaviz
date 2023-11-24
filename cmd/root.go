/*
Copyright © 2023 Simon Dahlbacka <simon.dahlbacka@fellowmind.fi>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "metaviz",
	Short: "Generate an ER diagram from metadata",
	Long: `
	Generate an ER diagram from metadata.

	The generator will generate a file in the ER format, which can be opened with the ER diagram tool of your choice (read: BurntSushi/erd).
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goerd.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
