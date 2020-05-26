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
	"fmt"
	"os"

	"github.com/BraspagDevelopers/bphc/lib"
	"github.com/spf13/cobra"
)

// healthCmd represents the health command
var healthCmd = &cobra.Command{
	Use:     "healthy",
	Aliases: []string{"h", "health", "healthcheck"},
	Short:   "Check if the site is healthy",
	Long: `Sends a GET HTTP request to a site in order to check its health.
	If the site returns a status code in the range 200-299 and the body is in JSON format and the value of the property IsHealthy is true, the site is considered healhty.
	If not, the check will fail.
	
	When the check succedes, it will be produce an exit code of 0. Any failure will produce a difference exit code. Additionaly, there will always be a message in STDOUT when the check fails`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		baseUrl := args[0]
		err := lib.HealthCheck(baseUrl, healthcheckPathFlag, verboseFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		fmt.Printf("The site %s is alive.\n", baseUrl)
	},
}

var healthcheckPathFlag string

func init() {
	rootCmd.AddCommand(healthCmd)
	healthCmd.Flags().StringVar(&healthcheckPathFlag, "path", "/healthcheck", "The path for the healthcheck endpoint")
}