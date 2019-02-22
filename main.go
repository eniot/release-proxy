package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	proxyCmd := proxyCmd()
	if err := proxyCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func proxyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proxy",
		Short: "Github Release Proxy For ESPHttpUpdate",
		Long:  "Github Release Proxy for ESPHttpUpdate",
		Run: func(cmd *cobra.Command, args []string) {
			flags := cmd.Flags()
			proxy(flags.Lookup("repo").Value.String(), flags.Lookup("addr").Value.String())
		},
	}
	cmd.Flags().StringP("repo", "r", "eniot/esp8266-firmware", "Release repo")
	cmd.Flags().StringP("addr", "a", ":8080", "API listening address")
	viper.BindPFlag("repo", cmd.Flags().Lookup("repo"))
	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	return cmd
}
