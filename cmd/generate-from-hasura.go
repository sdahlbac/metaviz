/*
Copyright © 2023 Simon Dahlbacka <simon.dahlbacka@fellowmind.fi>
*/
package cmd

import (
	"log"

	"github.com/sdahlbac/metaviz/internal/hasura_generator"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateFromHasuraCmd = &cobra.Command{
	Use:   "generate-from-hasura",
	Short: "Generate an ER diagram from hasura metadata",
	Long: `Generate an ER diagram from metadata.

	Example:
	metaviz generate-from-hasura --hasura-url foo --admin-secret bar --output ./erdiagram.er
	`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			log.Fatalf("Error getting output flag: %v", err)
		}
		includeProperties := cmd.Flag("properties").Value.String() == "true"
		hasuraUrl, err := cmd.Flags().GetString("hasura-url")
		if err != nil {
			log.Fatalf("Error getting hasura-url flag: %v", err)
		}
		adminSecret, err := cmd.Flags().GetString("admin-secret")
		if err != nil {
			log.Fatalf("Error getting admin-secret flag: %v", err)
		}
		hasura_generator.GenerateFromHasura(hasuraUrl, adminSecret, output, includeProperties)
	},
}

func init() {
	rootCmd.AddCommand(generateFromHasuraCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	generateFromHasuraCmd.Flags().StringP("output", "o", "model.er", "Output filename")
	generateFromHasuraCmd.PersistentFlags().BoolP("properties", "p", false, "Include properties in the diagram")
	generateFromHasuraCmd.Flags().StringP("hasura-url", "u", "", "Hasura URL")
	generateFromHasuraCmd.Flags().StringP("admin-secret", "s", "", "Hasura admin secret")
}
