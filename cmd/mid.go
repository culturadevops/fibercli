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

// midCmd represents the mid command
var midCmd = &cobra.Command{
	Use:   "mid",
	Short: "Crea archivos middlewares",
	Long: ` crea archivo de mid en y lo relaciona en el main 
pero si usa la bandera -p o --parent apuntara a otro archivo`,
	Run: func(cmd *cobra.Command, args []string) {
		if flags, _ := cmd.Flags().GetBool("parent"); flags {
			services.VarSrv.CreateMidDefaultToRoutes(services.Dirs.Mid, args[0], args[1], services.Dirs.Mid)
		} else {
			services.VarSrv.CreateMidDefault(services.Dirs.Mid, services.Dirs.Mid, args[0])
		}
	},
}

func init() {
	midCmd.Flags().BoolP("parent", "p", false, "Cambia el destino del mid a un routes, sino coloca este flag el mid se pondra en main")
	mkCmd.AddCommand(midCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// midCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// midCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
