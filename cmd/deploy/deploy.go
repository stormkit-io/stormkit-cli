package deploy

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/stormkit-io/stormkit-cli/api"
	"github.com/stormkit-io/stormkit-cli/cmd"
	"github.com/stormkit-io/stormkit-cli/model"
	"github.com/stormkit-io/stormkit-cli/stormkit"
	"github.com/stormkit-io/stormkit-cli/utils"
)

// interactiveFlag
const interactiveFlag = "interactive"

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy <env> <branch>",
	Short: "Initiate a deployment process",
	Long:  `Initiate a deployment process by providing the environment and branch names. The specified environment will be used to read the configuration to use while building the given branch.`,
	RunE:  runDeploy,
}

func init() {
	cmd.GetRootCmd().AddCommand(deployCmd)
	deployCmd.Flags().BoolP(interactiveFlag, "i", false, "Use command as interactive")
}

func runDeploy(cmd *cobra.Command, args []string) error {
	interactive, err := cmd.Flags().GetBool(interactiveFlag)
	if err != nil {
		return err
	}

	if interactive {
		envs, err := api.Envs(stormkit.GetEngineAppID())
		if err != nil {
			return err
		}

		envPrompt := promptui.Select{
			Label: "Select env",
			Items: envs.Names(),
		}

		envIndex, env, err := envPrompt.Run()
		if err != nil {
			return err
		}

		path, err := utils.GitRoot()
		if err != nil {
			return err
		}
		r, err := git.PlainOpen(path)
		if err != nil {
			return err
		}
		branches, err := utils.GitBranchesNames(r)
		if err != nil {
			return err
		}
		branchesS := []string{"default"}
		branchesS = append(branchesS, branches...)

		branchPrompt := promptui.SelectWithAdd{
			Label:    "Select deploy branch",
			Items:    branchesS,
			AddLabel: "Other",
		}

		branchIndex, branch, err := branchPrompt.Run()
		if err != nil {
			return err
		}
		if branchIndex == 0 {
			branch = envs.Envs[envIndex].Branch
		}

		d := model.Deploy{
			AppID:  stormkit.GetEngineAppID(),
			Env:    env,
			Branch: branch,
		}

		deploy, err := api.Deploy(d)
		if err != nil {
			return err
		}

		fmt.Printf("Deploy ID: %s\n", deploy.ID)
		return nil
	}

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
