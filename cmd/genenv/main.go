package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/marcozac/genenv"
)

var (
	cfg     *genenv.Config
	cfgFile string
	fatal   = log.Fatal
)

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "genenv",
		Short: "A generator for environment variables based configurations",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			cfg, err = genenv.ReadConfig(cfgFile)
			if err != nil {
				return fmt.Errorf("configuration file not found in %s", cfgFile)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return genenv.Generate(cfg)
		},
	}
	cmd.PersistentFlags().
		StringVar(&cfgFile, "config", ".genenv.yml", "configuration file path")
	return cmd
}

func main() {
	err := rootCmd().Execute()
	if err != nil {
		fatal(err)
	}
}
