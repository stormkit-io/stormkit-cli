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

// runDeploy executes the deploy command, it checks if use interactive or
// params mode for read the env and branch to deploy then calls api and prints the deploy ID
func runDeploy(cmd *cobra.Command, args []string) error {
	// read interactive flag
	interactive, err := cmd.Flags().GetBool(interactiveFlag)
	if err != nil {
		return err
	}

	// check if run interactive or params mode
	var d *model.Deploy
	if interactive {
		d, err = deployInteractive()
	} else {
		d, err = deployParams(args)
	}

	// check errors from interactive or params mode
	if err != nil {
		return err
	}

	// call stormkit API's for start deployment
	deploy, err := api.Deploy(*d)
	if err != nil {
		return err
	}

	// print deploy id
	fmt.Printf("Deploy ID: %s\n", deploy.ID)
	return nil
}

// deployParams creates the deployment struct via the cli params
func deployParams(args []string) (*model.Deploy, error) {
	// check if both args (env, branch) are present
	if len(args) < 2 {
		return nil, fmt.Errorf("not enought arguments")
	}

	// build struct
	return &model.Deploy{
		AppID:  stormkit.GetEngineAppID(),
		Env:    args[0],
		Branch: args[1],
	}, nil
}

// runEnvPrompt asks to chose between all the availables envs
func runEnvPrompt(envs *model.EnvsArray) (int, error) {
	// create envPrompt
	envPrompt := promptui.Select{
		Label: "Select env",
		Items: envs.Names(),
	}

	// run Env prompt
	i, _, err := envPrompt.Run()
	if err != nil {
		return -1, err
	}

	return i, nil
}

// runBranchPrompt check if command executed inside a git repository
// directory, if not ask the branch manually otherways ask to select
// a branch of the repository (allowed also default env branch, or
// other branch, only remote branch).
func runBranchPrompt(defaultBranch string) (string, error) {
	// get git repository root path
	path, err := utils.GitRoot()
	if err != nil {
		// if no git root path ask manually branch
		branchPrompt := promptui.Prompt{
			Label: "Branch",
		}
		return branchPrompt.Run()
	}

	// get repository
	r, err := git.PlainOpen(path)
	if err != nil {
		return "", err
	}

	// get git repository branches names
	branches, err := utils.GitBranchesNames(r)
	if err != nil {
		return "", err
	}

	// inserting "default" at index 0
	branches = append([]string{"default"}, branches...)

	// building branch prompt
	branchPrompt := promptui.SelectWithAdd{
		Label:    "Select branch to deploy",
		Items:    branches,
		AddLabel: "Other",
	}

	// running branch prompt
	branchIndex, branch, err := branchPrompt.Run()
	if err != nil {
		return "", err
	}
	// check if default branch
	if branchIndex == 0 {
		return defaultBranch, nil
	}

	return branch, nil
}

// deployInteractive creates the the deploy struct for do the api request
// for the deployment, it requests to chose the enviroment where to deploy
// and then the branch to deploy
func deployInteractive() (*model.Deploy, error) {
	// Get Envs via stormkit API's
	envs, err := api.Envs(stormkit.GetEngineAppID())
	if err != nil {
		return nil, err
	}

	// select env via prompt
	envIndex, err := runEnvPrompt(envs)
	if err != nil {
		return nil, err
	}

	// selct branch via prompt
	branch, err := runBranchPrompt(envs.Envs[envIndex].Branch)
	if err != nil {
		return nil, err
	}

	// build deploy struct
	return &model.Deploy{
		AppID:  stormkit.GetEngineAppID(),
		Env:    envs.Envs[envIndex].Env,
		Branch: branch,
	}, nil
}
