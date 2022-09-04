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
	"github.com/culturadevops/fibercli/services"
	"github.com/spf13/cobra"
)

// libCmd represents the lib command
var libCmd = &cobra.Command{
	Use:   "lib",
	Short: "Crea im archivo en libs ",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		flags, _ := cmd.Flags().GetBool("crud")
		if flags || services.Dirs.FlagCrud == "true" {

		}
		services.VarSrv.CreateLibDefault(services.Dirs.Libs, services.Dirs.Libs, args[0])
	},
}

func init() {
	mkCmd.AddCommand(libCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// libCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// libCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
