package deploy

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stormkit-io/stormkit-cli/api"
	"github.com/stormkit-io/stormkit-cli/cmd"
	"github.com/stormkit-io/stormkit-cli/model"
	"github.com/stormkit-io/stormkit-cli/stormkit"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Initiate a deployment process",
	Long:  `Initiate a deployment process by providing the environment and branch names. The specified environment will be used to read the configuration to use while building the given branch.`,
	RunE:  runDeploy,
}

func init() {
	cmd.GetRootCmd().AddCommand(deployCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runDeploy(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("not enought arguments")
	}

	d := model.Deploy{
		AppID:  stormkit.GetEngineAppID(),
		Env:    args[0],
		Branch: args[1],
	}

	deploy, err := api.Deploy(d)
	if err != nil {
		return err
	}

	fmt.Printf("Deploy ID: %s\n", deploy.ID)
	return nil
}
