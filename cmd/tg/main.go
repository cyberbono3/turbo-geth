package main

import (
	"fmt"
	"net"
	"os"

	"github.com/ledgerwatch/turbo-geth/cmd/utils"
	"github.com/ledgerwatch/turbo-geth/eth/stagedsync"
	"github.com/ledgerwatch/turbo-geth/log"
	turbocli "github.com/ledgerwatch/turbo-geth/turbo/cli"
	"github.com/ledgerwatch/turbo-geth/turbo/node"
	"github.com/urfave/cli"
)

var (
	// gitCommit is injected through the build flags (see Makefile)
	gitCommit string
)

func main() {
	// creating a turbo-api app with all defaults
	app := turbocli.MakeApp(runTurboGeth, turbocli.DefaultFlags)
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runTurboGeth(cliCtx *cli.Context) {
	// creating staged sync with all default parameters
	sync := stagedsync.New(
		stagedsync.DefaultStages(),
		stagedsync.DefaultUnwindOrder(),
		stagedsync.OptionalParameters{},
	)

	ctx := utils.RootContext()

	// initializing the node and providing the current git commit there
	tg := node.New(cliCtx, sync, node.Params{GitCommit: gitCommit})
	tg.SetP2PListenFunc(func(network, addr string) (net.Listener, error) {
		var lc net.ListenConfig
		return lc.Listen(ctx, network, addr)
	})
	// running the node
	err := tg.Serve()

	if err != nil {
		log.Error("error while serving a turbo-geth node", "err", err)
	}
}
