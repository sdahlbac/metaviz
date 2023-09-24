/*
Copyright © 2023 Simon Dahlbacka <simon.dahlbacka@fellowmind.fi>
*/
package cmd

import (
	"log"

	"github.com/sdahlbac/metaviz/internal/generator"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate an ER diagram from metadata",
	Long: `Generate an ER diagram from metadata.

	Example:
	metaviz generate ./metadata --output ./erdiagram.er
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			log.Fatalf("Error getting output flag: %v", err)
		}
		includeProperties := cmd.Flag("properties").Value.String() == "true"
		generator.Generate(args[0], output, includeProperties)
	},
}

var properties bool

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	generateCmd.Flags().StringP("output", "o", "model.er", "Output filename")
	generateCmd.PersistentFlags().BoolP("properties", "p", false, "Include properties in the diagram")
}
