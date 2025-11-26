package commands

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use: "atlans",
}

func Run(args []string) error {
	RootCmd.SetArgs(args)
	return RootCmd.Execute()
}
