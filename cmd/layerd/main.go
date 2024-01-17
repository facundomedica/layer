package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/tellor-io/layer/app"
	"github.com/tellor-io/layer/cmd/layerd/cmd"
)

func main() {
	option := cmd.GetOptionWithCustomStartCmd()
	rootCmd, _ := cmd.NewRootCmd(option)
	cmd.AddInitCmdPostRunE(rootCmd)
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
