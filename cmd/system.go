/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/syxaxis/automicv2/pkg"
)

// systemCmd represents the system command
var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "Get info on the system",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("system called")

		dispType, _ := cmd.Flags().GetString("type")
		showJobsOnly, _ := cmd.Flags().GetBool("showjobsonly")

		switch dispType {
		case "system":
			pkg.GetSystemStatus(globalParams)
		case "active":
			pkg.GetActiveExecutions(showJobsOnly, globalParams)
		case "agents":
			pkg.GetSystemAgents(globalParams)
		}

	},
}

func init() {
	rootCmd.AddCommand(systemCmd)

	systemCmd.Flags().StringP("type", "t", "system", "system, active, agents")
	systemCmd.Flags().BoolP("showjobsonly", "j", false, "show jobs, workflows and scripts only")

}
