/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/syxaxis/automicv2/pkg"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export objects",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("export called")

		objectName, _ := cmd.Flags().GetString("objectname")
		exportfile, _ := cmd.Flags().GetString("exportfile")
		showjson, _ := cmd.Flags().GetBool("showjson")

		responseData, err := pkg.GetExportObjectCode(objectName, globalParams)
		if err != nil {
			os.Exit(1)
		}

		switch showjson {
		case true:
			var jsonData pkg.SingleObjectDump
			json.Unmarshal(responseData, &jsonData)
			// fmt.Println(string(jsonData.RawJSONData))
			rawjson, _ := json.MarshalIndent(jsonData.RawJSONData, "", "  ")
			fmt.Println(string(rawjson))
			fmt.Printf("Folder path : %v\n", jsonData.FolderPath)

		case false:
			err = os.WriteFile(exportfile, []byte(responseData), 0644)
			if err != nil {
				log.Fatalln(err)
				os.Exit(1)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	exportCmd.Flags().StringP("objectname", "n", "TESTSCHED_WIN", "Name of single object to export")
	exportCmd.Flags().StringP("exportfile", "f", "export.json", "Filename to write exported object spec to")
	exportCmd.Flags().BoolP("showjson", "s", false, "simply display the JSON def")

}
