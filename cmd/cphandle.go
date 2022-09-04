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
	"github.com/atotto/clipboard"
	"github.com/culturadevops/fibercli/services"
	"github.com/spf13/cobra"
)

// cphdeCmd represents the cphde command
var handleCmd = &cobra.Command{
	Use:   "hde",
	Short: "Copia una funcion handle en el Portapapeles",
	Long: ` Comando cp hde nombredefuncion creara una funcion en memoria que podras pegar;  
La funcion tiene el siguiente formato: 
func nombredefuncion() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if u == nil {
			return c.Status(http.StatusOK).JSON(data)
		}
			return c.Status(http.StatusInternalServerError).JSON(err)
	}
}

Estas funciones son llamados handle y sirve para resolver la 
peticiones capturada por el servidor de fiber
	`,
	Run: func(cmd *cobra.Command, args []string) {
		texto := services.VarSrv.GetCtlCrud(services.Dirs.Hde, services.Dirs.Hde, args[0])
		clipboard.WriteAll(texto)
	},
}

func init() {
	cpCmd.AddCommand(handleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cphdeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cphdeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
