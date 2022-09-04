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

// rtsCmd represents the rts command
var rtsCmd = &cobra.Command{
	Use:   "rts",
	Short: "Crea archivos routes",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if flags, _ := cmd.Flags().GetBool("api"); flags {
			services.VarSrv.CreateRtsApi(services.Dirs.Rts, args[0], services.Dirs.Rts)
		} else {
			services.VarSrv.CreateRtsDefault(services.Dirs.Rts, args[0], services.Dirs.Rts)
		}

	},
}

func init() {
	rtsCmd.Flags().BoolP("api", "a", false, "crea una plantilla crud")
	mkCmd.AddCommand(rtsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rtsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rtsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
