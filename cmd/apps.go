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
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/giuliobosco/stormcli/model"
)

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
	appsCmd.Flags().BoolP("details", "d", false, "Show details of the apps")
	appsCmd.Flags().BoolP("numbers", "n", false, "Show the index numbers of the applications")
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
	bearerToken := viper.GetString("app.bearer_token")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
	
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

	details, err := cmd.Flags().GetBool("details")
	if err != nil {
		log.Fatalln(err)
	}
	numbers, err := cmd.Flags().GetBool("numbers")
	if err != nil {
		log.Fatalln(err)
	}

	printf := "%" + strconv.Itoa(len(apps.Apps)-1) + "v %s\n"
	tabf := fmt.Sprintf("%" + strconv.Itoa(len(apps.Apps)-1) + "v    ", "")
	for i, a := range apps.Apps {
		if numbers {
			//fmt.Printf("%d | %s\n", i, a.Repo)
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
