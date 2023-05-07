/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/syxaxis/automicv2/pkg"
)

var globalParams pkg.GlobalParams

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "automicv2",
	Short: "short desc TBC",
	Long:  "long desc TBC",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	// NOTE - flags get parsed after this call
	//        you can tap the values form the global flags below from the var rootParams

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aut01.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVar(&globalParams.AEHostname, "aehost", "automicAEserver", "AE hostname")
	rootCmd.PersistentFlags().StringVar(&globalParams.AEPort, "aeport", "8088", "AE host REST port")
	rootCmd.PersistentFlags().StringVar(&globalParams.AEClientID, "aeclientid", "333", "AE client ID")

	// aehost, _ := rootCmd.Flags().GetString("aehost")
	// aeport, _ := rootCmd.Flags().GetString("aeport")
	// aeclientid, _ := rootCmd.Flags().GetString("aeclientid")

}
