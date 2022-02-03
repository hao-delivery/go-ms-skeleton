package main

import (
	"github.com/spf13/cobra"

	"github.com/hao-delivery/go-ms-skeleton/cmd"
)

func main() {
	cobra.CheckErr(cmd.NewRootCmd().Execute())
}
