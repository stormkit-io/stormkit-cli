package log

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stormkit-io/stormkit-cli/api"
	"github.com/stormkit-io/stormkit-cli/cmd"
	"github.com/stormkit-io/stormkit-cli/stormkit"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Log of deploy",
	Long:  `Print the logs of the deployment through his ID and the ID of the application, the id of the application from file ~/.stomkit-cli.yml or ./stormkit.config.yml`,
	RunE:  runLog,
}

func init() {
	cmd.GetRootCmd().AddCommand(logCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runLog(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("not enought arguments")
	}

	appID := stormkit.GetEngineAppID()
	id := args[0]

	deploy, err := api.DeployByID(appID, id)
	if err != nil {
		return err
	}

	for _, l := range deploy.Deploy.Logs {
		fmt.Println(l.Dump())
	}

	return nil
}
