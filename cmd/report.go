/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/syxaxis/automicv2/pkg"
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Report status of task execs",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("report called")

		taskRunID, _ := cmd.Flags().GetInt("runid")
		reportType, _ := cmd.Flags().GetString("reporttype")

		searchitemname, _ := cmd.Flags().GetString("searchitemname")
		searchitemtype, _ := cmd.Flags().GetString("searchitemtype")
		searchStartTime, _ := cmd.Flags().GetString("searchstarttime")
		searchEndTime, _ := cmd.Flags().GetString("searchendtime")
		searchDeactivated, _ := cmd.Flags().GetBool("searchincdeactivated")

		switch strings.ToUpper(reportType) {
		case "LIST":
			pkg.GetPreviousExecutions(searchitemname, searchitemtype, searchStartTime, searchEndTime, searchDeactivated, globalParams)
		case "REP", "ACT", "LOG", "POST":
			pkg.GetTaskReportOutput(taskRunID, strings.ToUpper(reportType), globalParams)
		case "CHILD":
			pkg.GetChildExecutions(taskRunID, globalParams)
		default:
			pkg.GetTaskReportsAvailable(taskRunID, globalParams)
		}
		os.Exit(0)

	},
}

func init() {
	rootCmd.AddCommand(reportCmd)

	reportCmd.Flags().IntP("runid", "r", 8749574, "Task RunID ( obtain from [ display --type active ] )")
	reportCmd.Flags().String("reporttype", "REP", "Report type to show [LIST or REP,ACT,LOG,POST or CHILD] ")

	reportCmd.Flags().String("searchitemname", "*", "Item name")
	reportCmd.Flags().String("searchitemtype", "JOBS", "Search type : eg, JOBS, SCRI, JOBP, etc")
	reportCmd.Flags().StringP("searchstarttime", "s", time.Now().Add(-time.Hour*24).Format("2006-01-02T15:04:05Z"), "search range start time")
	reportCmd.Flags().StringP("searchendtime", "e", time.Now().Format("2006-01-02T15:04:05Z"), "search range end time")
	reportCmd.Flags().BoolP("searchincdeactivated", "d", true, "show deactivated tasks")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
