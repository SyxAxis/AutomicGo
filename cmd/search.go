/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/syxaxis/automicv2/pkg"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for objects",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")

		// srchType, _ := cmd.Flags().GetString("itemtype")
		objSrchName, _ := cmd.Flags().GetString("objectname")
		objSrchMaxRslt, _ := cmd.Flags().GetInt("maxsrchresults")
		objSrchType, _ := cmd.Flags().GetStringArray("objecttypes")

		pkg.SearchForObjectsByName(objSrchName, objSrchMaxRslt, objSrchType, globalParams)

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringP("objectname", "n", "GDF_TASK_0870_GRRT_PTT_TO_SCD_TRAX", "Item name")
	searchCmd.Flags().IntP("maxsrchresults", "m", 10, "Number of results to return from a search")

	searchCmd.Flags().StringArrayP("objecttypes", "t", []string{"JOBS"}, "Search type : eg, JOBS,JOBP or SCRI,FOLD,EVNT or simply \"*\"")

}
