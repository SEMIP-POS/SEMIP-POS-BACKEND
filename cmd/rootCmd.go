package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile  string
	scrtFile string
)

var rootCmd = &cobra.Command{
	Use:   "cobra-cli",
	Short: "A generator for Simple Cobra App",
	Long: `Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "")
	rootCmd.PersistentFlags().StringVar(&scrtFile, "secret", "secret.yaml", "")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
