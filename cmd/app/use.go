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

	"github.com/spf13/cobra"
	"github.com/stormkit-io/stormkit-cli/api"
	"github.com/stormkit-io/stormkit-cli/stormkit"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Select an application to work on",
	Long: `Select an application where to work on, via his repo name, id or index.
Via his repo name:

$ stormkit-cli app use <repo_name>

Via his id:

$ stormkit-cli app use --app-id <app_id>

Via his index:

$ stormkit-cli app ls -n
0 github/myuser/my-project
$ stormkit-cli app use 0

This command is useful when you are not in the directory of your repository.
It keeps his memory in the config file, at: app.engine.app_id.
Be carefull this wouldn't work if you are in the direcotry of another 
stormkit project.`,
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

	// using index
	stormkit.SetEngineAppID(apps.Apps[index].ID)
	return nil
}
