package main

import (
	"fmt"
	"os"

	"github.com/ansrivas/getignore/downloader"
	"github.com/spf13/cobra"
)

var (
	Downloader = downloader.New()
	rootCmd    = &cobra.Command{Use: "getignore"}
)

func init() {

	rootCmd.AddCommand(cmdList)
	rootCmd.AddCommand(cmdDwnld)

}

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "Display a list of available gitignore files.",
	Long:  `Display a list of available gitignore files available on github.`,
	Run: func(cmd *cobra.Command, args []string) {
		Downloader.ListLanguages(true)
	},
}

var cmdDwnld = &cobra.Command{
	Use:     "download [language to download gitignore for]",
	Short:   "Download a gitignore file for the given language.",
	Long:    `Download a gitignore from github for a given language.`,
	Example: "getignore download python",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			curcmd, _, _ := rootCmd.Find([]string{"download"})
			fmt.Printf("\033[31m%s\033[39m\n\n", curcmd.UsageString())
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		Downloader.DownloadFile(args[0])
	},
}
