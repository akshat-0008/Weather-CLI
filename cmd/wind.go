/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"WeatherCli/pkg"

	"github.com/spf13/cobra"
)

// windCmd represents the wind command
var windCmd = &cobra.Command{
	Use:   "wind",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			pkg.GetWindInfo(args[0])
		} else {
			fmt.Println("Enter city name with command")
		}
	},
}

func init() {
	rootCmd.AddCommand(windCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// windCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// windCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
