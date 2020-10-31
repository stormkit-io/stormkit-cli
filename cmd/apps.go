package cmd
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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/giuliobosco/stormcli/model"
)

const authBearer = ""

// appsCmd represents the apps command
var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	/*Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("apps called")
	},*/
	Run: runApps,
}

func init() {
	rootCmd.AddCommand(appsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// appsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// appsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type AppsResponse struct {
	Apps []model.App `json:"apps"`
}

func runApps(cmd *cobra.Command, args []string) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("GET", fmt.Sprintf("https://api.stormkit.io/apps"), nil)
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authBearer))
	
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var apps AppsResponse
	err = json.Unmarshal(body, &apps)

	if err != nil {
		log.Fatalln(err)
	}

	for _, a := range apps.Apps {
		fmt.Println(a.Repo)
	}
}