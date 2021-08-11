package cmd

import (
	"fmt"
	"os"
	"github.com/Volatus/kdfn/pkg"
	"github.com/spf13/cobra"
)

var (
	label string
)

var rootCmd = &cobra.Command{
	Use:   "kdfn",
	Short: "kdfn is a kubectl plugin for checking disk usage of K8s nodes",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at http://hugo.spf13.com`,
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.ListNodes(pkg.GetNodes(label))
	},
}

func init() {
	rootCmd.Flags().StringVarP(&label, "", "l", "", "Labels to filter nodes: label=value")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
