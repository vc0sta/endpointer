/*
Copyright © 2020 NAME HERE vinicius.costa.92@gmail.com

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
	"github.com/d1m3/endpointer/cmd/check"
	// "github.com/d1m3/endpointer/cmd/check/postgres"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var Cmd = &cobra.Command{
	Use:   "check",
	Short: "Checks if a given resource is responding correctly",
	Long:  `You can check if any of the given resources are responding correctly`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {

	Cmd.AddCommand(check.PostgresCmd, check.HttpCmd, check.MYSQLCmd)
	// Cmd.AddCommand(check.HttpCmd)
	rootCmd.AddCommand(Cmd)
}
