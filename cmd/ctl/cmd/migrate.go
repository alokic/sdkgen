// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	// migrateCmd represents the migrate command
	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migrate a postgres database.",
		Long: `This command migrates a postgres database. 
		For example:

	cli migrate --db_url <postgres_db_url>`,
		Run: func(cmd *cobra.Command, args []string) {
			initLogger()
			initDB(dbURL)

			commmand := exec.Command("goose", "-dir", "app/postgres/migrate", "postgres", dbURL, "up")
			err := commmand.Run()
			if err != nil {
				fmt.Printf("migratioin error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("migratioin success")
		},
	}

	migrateDBURL string
)

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	migrateCmd.Flags().StringVar(&dbURL, "db_url", "", "postgres database url (required)")
	migrateCmd.MarkFlagRequired("db_url")
}
