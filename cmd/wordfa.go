/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"CiFa/cliserve"
	"CiFa/util/sortalgo"
	"CiFa/util/strsearch"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var wordfaCliServe = cliserve.CliWordfaServer{}

// wordfaCmd represents the wordfa command
var wordfaCmd = &cobra.Command{
	Use:   "wordfa",
	Short: "Run a words frequency analyzing task in CLI",
	Long:  `Run a words frequency analyzing task in CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		if wordfaCliServe.KeywordFilePath == "" || wordfaCliServe.SourceFilePath == "" {
			fmt.Println("Cannot run without KeywordFilePath & SourceFilePath given.")
			os.Exit(1)
		}
		fmt.Println("wordfa calling...")
		wordfaCliServe.Run()
	},
}

func init() {
	rootCmd.AddCommand(wordfaCmd)

	wordfaCmd.Flags().StringVarP(
		&wordfaCliServe.KeywordFilePath,
		"keyword", "k", "", "keywords file `path`",
	)
	wordfaCmd.Flags().StringVarP(
		&wordfaCliServe.SourceFilePath,
		"file", "f", "", "source file/dir `path`",
	)

	matchAlgorithmsName := ""
	for k, _ := range strsearch.StrsearchAlgorithmsMap {
		matchAlgorithmsName += k + ", "
	}
	wordfaCmd.Flags().StringVarP(
		&wordfaCliServe.StrsearchAlgo,
		"match", "m", "",
		"string match `algorithm`: one of "+strings.Trim(matchAlgorithmsName, ", "),
	)

	sortAlgorithmsName := ""
	for k, _ := range sortalgo.SortAlgorithmsMap {
		sortAlgorithmsName += k + ", "
	}
	wordfaCmd.Flags().StringVarP(
		&wordfaCliServe.StrsearchAlgo,
		"sort", "s", "",
		"result sort `algorithm`: one of "+strings.Trim(sortAlgorithmsName, ", "),
	)

	wordfaCmd.Flags().StringVarP(
		&wordfaCliServe.OutputFilePath,
		"output", "o", "", "output result to `file`",
	)
}
