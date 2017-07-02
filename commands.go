// MIT License
//
// Copyright (c) 2017 Ankur Srivastava
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"fmt"
	"os"

	"github.com/ansrivas/getignore/downloader"
	"github.com/spf13/cobra"
)

var (
	getIgnore  = downloader.NewGitIgnore()
	getLicense = downloader.NewLicense()
	rootCmd    = &cobra.Command{Use: "getignore"}
)

func init() {

	rootCmd.AddCommand(cmdIgnoreList)
	rootCmd.AddCommand(cmdIgnoreDwnld)
	rootCmd.AddCommand(cmdLicenseList)
	rootCmd.AddCommand(cmdLicenseDwnld)

}

var cmdIgnoreList = &cobra.Command{
	Use:   "listIgnores",
	Short: "Display a list of available gitignore files.",
	Long:  `Display a list of available gitignore files available on github.`,
	Run: func(cmd *cobra.Command, args []string) {
		getIgnore.ListLanguages(true)
	},
}

var cmdLicenseList = &cobra.Command{
	Use:   "listLicenses",
	Short: "Display a list of available licenses.",
	Long:  `Display a list of available license files available on github.`,
	Run: func(cmd *cobra.Command, args []string) {
		getLicense.ListLicenses(true)
	},
}

var cmdIgnoreDwnld = &cobra.Command{
	Use:     "downloadIgnore [language to download gitignore for]",
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
		getIgnore.DownloadFile(args[0])
	},
}

var cmdLicenseDwnld = &cobra.Command{
	Use:     "downloadLicense [license to download]",
	Short:   "Download a license file from github.",
	Long:    `Download a license file from github for a given license string.`,
	Example: "getignore downloadLicense lgpl-3.0",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			curcmd, _, _ := rootCmd.Find([]string{"downloadLicense"})
			fmt.Printf("\033[31m%s\033[39m\n\n", curcmd.UsageString())
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		getLicense.DownloadFile(args[0])
	},
}
