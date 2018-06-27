package main

import (
	"github.com/koshatul/registry-cli/src/registry"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test Command",
	Run:   testCommand,
}

func init() {
	configDefaults()

	rootCmd.AddCommand(testCmd)
}

func testCommand(cmd *cobra.Command, args []string) {
	c, err := registry.Dial(
		viper.GetString("registry.url"),
		viper.GetString("registry.username"),
		viper.GetString("registry.password"),
	)
	if err != nil {
		logrus.Fatal(err)
	}
	repos, err := c.ListRepositories()
	if err != nil {
		logrus.Error(err)
	}

	for _, repo := range repos {
		logger := logrus.WithField("repo", repo)
		logger.Infof("Repo: %s", repo)
		tags, err := c.ListTags(repo)
		if err != nil {
			logrus.Error(err)
			continue
		}
		for _, tag := range tags {
			logger = logger.WithField("tag", tag)
			logger.Infof("Tag: %s", tag)
		}
	}

}
