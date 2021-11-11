/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gatherCmd represents the gather2 command
var gatherCmd = &cobra.Command{
	Use:   "gather",
	Short: "Gathers knowledge from wikipedia (DANGEROUS)",
	Long: `Not really dangerous. But useless once wac has been compiled.
The gather command is used at build time of wac. It fetches information
from https://en.wikipedia.org/wiki/Country_code_top-level_domain and
puts it into Go source files, which become part of the final wac binary.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gather called")
	},
}

func init() {
	dbCmd.AddCommand(gatherCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gather2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gather2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
