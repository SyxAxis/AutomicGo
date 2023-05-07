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

// displayCmd represents the display command
var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "Display info from the system",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("display called")

		dispType, _ := cmd.Flags().GetString("type")
		showJobsOnly, _ := cmd.Flags().GetBool("showjobsonly")
		taskRunID, _ := cmd.Flags().GetInt("runid")

		switch dispType {
		case "system":
			pkg.GetSystemStatus(globalParams)
			os.Exit(0)
		case "active":
			pkg.GetActiveExecutions(showJobsOnly, globalParams)
			os.Exit(0)
		case "jobmeta":
			pkg.GetActiveRuntimeMetadata(taskRunID, globalParams)
			os.Exit(0)
		}

	},
}

func init() {
	rootCmd.AddCommand(displayCmd)

	displayCmd.Flags().StringP("type", "t", "system", "system, active, ...")
	displayCmd.Flags().BoolP("showjobsonly", "j", false, "show jobs, workflows and scripts only")
	displayCmd.Flags().IntP("runid", "r", 8749574, "Task RunID ( obtain from [ display --type active ] )")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// displayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// displayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
