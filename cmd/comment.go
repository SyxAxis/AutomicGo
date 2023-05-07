/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/syxaxis/automicv2/pkg"
)

// commentCmd represents the comment command
var commentCmd = &cobra.Command{
	Use:   "comment",
	Short: "Add remove comments on tasks",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("comment called")

		runid, _ := cmd.Flags().GetInt("runid")
		addcomment, _ := cmd.Flags().GetString("add")
		if len(addcomment) == 0 {
			pkg.GetTaskComments(runid, globalParams)
		} else {
			pkg.SetTaskComment(runid, addcomment, globalParams)
		}

	},
}

func init() {
	rootCmd.AddCommand(commentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	commentCmd.Flags().IntP("runid", "r", 8749574, "Task RunID ( obtain from [ display --type active ] )")
	commentCmd.Flags().StringP("add", "a", "", "Add new comment ( of ommitted simply reads current comments )")

}
