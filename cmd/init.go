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
	"github.com/culturadevops/fibercli/services"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Inicializa un proyecto",
	Long: `Inicializa un proyecto copiando una estructura base
que tiene lo minimo necesario para comenzar con tu proyecto de fiber
tienes la opcion MVC Y CLEARCODE`,
	Run: func(cmd *cobra.Command, args []string) {
		if flags, _ := cmd.Flags().GetBool("serverless"); flags {

			services.VarSrv.InitServerlessCopyAll(args[0])
		} else {
			if flags, _ := cmd.Flags().GetBool("cleancode"); flags {
				services.VarSrv.InitCleanCodeCopyAll(args[0])
			} else {
				services.VarSrv.InitMVCCopyAll(args[0])
			}
		}

	},
}

func init() {

	initCmd.Flags().BoolP("cleancode", "c", false, "crea un proyecto en cleancode")
	initCmd.Flags().BoolP("serverless", "s", false, "crea un proyecto en serverless")

	rootCmd.AddCommand(initCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
