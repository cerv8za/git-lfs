package commands

import "github.com/github/git-lfs/vendor/_nuts/github.com/spf13/cobra"

// TODO: Remove for Git LFS v2.0 https://github.com/github/git-lfs/issues/839

var (
	initCmd = &cobra.Command{
		Use:        "init",
		Deprecated: "Use 'git lfs install' now",
		Run:        initCommand,
	}

	initHooksCmd = &cobra.Command{
		Use:        "hooks",
		Deprecated: "Use 'git lfs install' now",
		Run:        initHooksCommand,
	}
)

func initCommand(cmd *cobra.Command, args []string) {
	installCommand(cmd, args)
}

func initHooksCommand(cmd *cobra.Command, args []string) {
	installHooksCommand(cmd, args)
}

func init() {
	initCmd.Flags().BoolVarP(&forceInstall, "force", "f", false, "Set the Git LFS global config, overwriting previous values.")
	initCmd.Flags().BoolVarP(&localInstall, "local", "l", false, "Set the Git LFS config for the local Git repository only.")
	initCmd.Flags().BoolVarP(&skipSmudgeInstall, "skip-smudge", "s", false, "Skip automatic downloading of objects on clone or pull.")
	initCmd.AddCommand(initHooksCmd)
	RootCmd.AddCommand(initCmd)
}
