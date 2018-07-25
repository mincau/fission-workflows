package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var cmdStatus = cli.Command{
	Name:    "status",
	Aliases: []string{"s"},
	Usage:   "Check cluster status",
	Action: commandContext(func(ctx Context) error {
		client := getClient(ctx)
		resp, err := client.Admin.Status(ctx)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf(resp.Status)

		return nil
	}),
}
