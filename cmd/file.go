/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	"github.com/spf13/cobra"
)

var output string
var version string
var noinput string
var output1 string

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use: "file",
	Short: ` yaml file to be merged, you may specify as many yaml files as you want. 
		Files should be ordered in ascending level of importance in the hierarchy. 
		A yaml node in the last file replaces values in any previous file. 
		If a text is piped to yaml_merger.py it's used as the first yaml to be merged
		`,
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("files to merge:", args[0], args[1])

		output1, _ := rootCmd.Flags().GetString("output")
		fmt.Println("Output:", output1)

		fmt.Println("Noinput:", noinput)
		/*
			noinput, _ := rootCmd.Flags().GetString("noinput")
			output, _ := rootCmd.Flags().GetString("output")
			if noinput != "" {
				fmt.Println("Files to merge:", args[0], args[1], ". No input flag active")
			}

			if output != "" {
				fmt.Println("Files to merge:", args[0], args[1], ". Output file", args[2])
			}
		*/
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	/*fileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	fileCmd.Flags().StringVar(&output, "output", "o", "write merged yaml to output file instead of stdout")
	fileCmd.Flags().StringVar(&version, "version", "v", "show program's version number and exit") */
	// TODO DE TIPO BOOLEAN
	fileCmd.Flags().StringVarP(&noinput, "no-input", "", "", "ignore stdin input (by default stdin is first file)")

}
