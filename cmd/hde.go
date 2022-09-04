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

	"github.com/culturadevops/fibercli/services"
	"github.com/spf13/cobra"
)

// hdeCmd represents the hde command
var hdeCmd = &cobra.Command{
	Use:   "hde",
	Short: "Crea un archivo handle en la carpetas de handles",
	Long: `Puedes crear un archivo de handle siguiente un patron crud 
usando la bandera -c o --crud`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("creando archivo modelo en la ruta ->" + services.Dirs.Mdl)
		flags, _ := cmd.Flags().GetBool("crud")
		if flags || services.Dirs.FlagCrud == "true" {
			services.VarSrv.CreateCtlCrud(services.Dirs.Hde, services.Dirs.Hde, args[0])
		} else {

			services.VarSrv.CreateCtlDefault(services.Dirs.Hde, services.Dirs.Hde, args[0])
		}
	},
}

func init() {
	hdeCmd.Flags().BoolP("crud", "c", false, "crea una plantilla crud")
	mkCmd.AddCommand(hdeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hdeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hdeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
