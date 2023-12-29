package cmd

import (
	"cobra-api/cmd/api"
	"cobra-api/cmd/cron"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "cobra-api",
	Short: "Go Blueprint application",
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start and run REST API",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting REST api...")
		api.NewApi()
	},
}

var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "Start scheduled cron jobs",
	Run: func(cmd *cobra.Command, args []string) {
		cron.NewCron()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("Error running app, Err: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(apiCmd)
	rootCmd.AddCommand(cronCmd)
}
