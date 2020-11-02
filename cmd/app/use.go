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
	"errors"
	"strconv"

	"github.com/stormkit-io/stormkit-cli/api"
	"github.com/stormkit-io/stormkit-cli/stormkit"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: runAppUse,
}

func init() {
	appCmd.AddCommand(useCmd)

	useCmd.Flags().StringP("app-id", "a", "", "ID of the app to use")
}

func runAppUse(cmd *cobra.Command, args []string) error {
	// Fetch Apps from  API
	apps, err := api.GetApps()
	if err != nil {
		return err
	}

	// Check if using app-id flag
	appID, err := cmd.Flags().GetString("app-id")
	if err != nil {
		return err
	}
	if len(appID) > 0 {
		for _, a := range apps.Apps {
			if a.ID == appID {
				stormkit.SetEngineAppID(a.ID)
				return nil
			}
		}

		return errors.New("no app found")
	}

	// check if arguments are present
	if len(args) == 0 {
		return errors.New("not enought arguments")
	}

	// using repo name or index
	index, err := strconv.Atoi(args[0])
	if err != nil {
			for _, a := range apps.Apps {
			if a.Repo == args[0] {
				stormkit.SetEngineAppID(a.ID)
				return nil
			}
		}

		return errors.New("app not found")
	}
	
	// checking index
	if index >= len(apps.Apps) {
		return errors.New("index too large")
	}

	stormkit.SetEngineAppID(apps.Apps[index].ID)
	return nil
}
