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

// middleCmd represents the middle command
var middleCmd = &cobra.Command{
	Use:   "mid",
	Short: "Copia una funciones middleware en el portapapeles",
	Long: `Comando cp mid Nombredemiddleware 
Crear una funciones similares a la de handles
pero esta diseñada para ser usada como middleware.
Ejemplo:
func  Nombredemiddleware(c *fiber.Ctx) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}

	`,
	Run: func(cmd *cobra.Command, args []string) {
		texto := services.VarSrv.CopyMidCrud(services.Dirs.Hde, services.Dirs.Hde, args[0])
		clipboard.WriteAll(texto)
	},
}

func init() {
	cpCmd.AddCommand(middleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// middleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// middleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
