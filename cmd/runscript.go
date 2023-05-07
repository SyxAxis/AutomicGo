/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/syxaxis/automicv2/pkg"
)

// runscriptCmd represents the runscript command
var runscriptCmd = &cobra.Command{
	Use:   "runscript",
	Short: "Run ad-hoc SCRI text",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("runscript called")

		jclScriptText, _ := cmd.Flags().GetString("scriptcode")

		fmt.Printf("[%v]\n", jclScriptText)

		pkg.RunScriptCode(jclScriptText, globalParams)
		os.Exit(0)

	},
}

func init() {
	rootCmd.AddCommand(runscriptCmd)

	runscriptCmd.Flags().StringP("scriptcode", "s", ":PRINT 'Hello from GWJ API CALL!'\\n:PRINT 'NEW line'", "Automic JCL script code")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runscriptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runscriptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
