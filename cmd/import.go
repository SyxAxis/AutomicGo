/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/syxaxis/automicv2/pkg"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("import called")

		importfile, _ := cmd.Flags().GetString("importfile")
		overwrite, _ := cmd.Flags().GetBool("overwrite")

		pkg.ImportObjectCode(importfile, overwrite, globalParams)

	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().StringP("importfile", "f", "import.json", "Filename to read exported object spec from")
	importCmd.Flags().BoolP("overwrite", "o", false, "Overwrite existing object is exists")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
