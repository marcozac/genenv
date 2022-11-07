package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/marcozac/genenv"
)

var (
	cfg     *genenv.Config
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "genenv",
		Short: "A generator for environment variables based configurations",
		PreRun: func(cmd *cobra.Command, args []string) {
			var err error
			cfg, err = genenv.ReadConfig(cfgFile)
			if err != nil {
				log.Fatalf("configuration file not found in %s", cfgFile)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			err := genenv.Generate(cfg)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func main() {
	rootCmd.PersistentFlags().
		StringVar(&cfgFile, "config", ".genenv.yml", "configuration file path")

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
