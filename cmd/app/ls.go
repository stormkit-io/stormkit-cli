package app

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

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/stormkit-io/stormkit-cli/api"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runAppLs,
}

func init() {
	appCmd.AddCommand(lsCmd)

	lsCmd.Flags().BoolP("details", "d", false, "Show details of the apps")
	lsCmd.Flags().BoolP("numbers", "n", false, "Show the index numbers of the applications")
}

func runAppLs(cmd *cobra.Command, args []string) {
	apps, err := api.GetApps()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	details, err := cmd.Flags().GetBool("details")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	numbers, err := cmd.Flags().GetBool("numbers")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lenf := strconv.Itoa(len(apps.Apps) - 1)
	printf := "%" + lenf + "v %s\n"
	tabf := fmt.Sprintf("%"+lenf+"v    ", "")

	for i, a := range apps.Apps {
		if numbers {
			fmt.Printf(printf, i, a.Repo)
		} else {
			fmt.Println(a.Repo)
		}

		if details {
			fmt.Printf("%sStatus: %t\n", tabf, a.Status)
			fmt.Printf("%sAutoDeploy: %s\n", tabf, a.AutoDeploy)
			fmt.Printf("%sDefaultEnv: %s\n", tabf, a.DefaultEnv)
			fmt.Printf("%sEndpoint: %s\n", tabf, a.Endpoint)
			fmt.Printf("%sDisplayName: %s\n", tabf, a.DisplayName)
			fmt.Printf("%sCreatedAt: %s\n", tabf, time.Unix(int64(a.CreatedAt), 0))
			fmt.Printf("%sDeployedAt: %s\n", tabf, time.Unix(int64(a.DeployedAt), 0))
			fmt.Println()
		}
	}
}
