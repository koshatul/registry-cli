package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var version = "dev"
var builddate = "notset"
var gittag = ""

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Run:   versionCommand,
}

func versionCommand(cmd *cobra.Command, args []string) {
	logrus.Debug("versionCommand():start")

	fmt.Printf("%s version %s [%s] (%s)\n", rootCmd.Name(), version, gittag, builddate)

	logrus.Debug("versionCommand():end")
}
