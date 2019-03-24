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

	"github.com/honestbee/authorizer/app/postgres"
	"github.com/spf13/cobra"
)

var (
	seedDBURL string

	// seedCmd represents the seed command
	seedCmd = &cobra.Command{
		Use:   "seed",
		Short: "Seed a postgres database.",
		Long: `This command seeds a postgres database. 
			For example:

		cli seed --db_url <postgres_db_url>`,
		Run: func(cmd *cobra.Command, args []string) {
			initLogger()
			initDB(dbURL)

			postgres.Clear(postgresDB, logger)
			postgres.Seed(postgresDB, fmt.Sprintf("%s/app/postgres/seed", projectRoot), logger)
			fmt.Println("seed success")
		},
	}
)

func init() {
	rootCmd.AddCommand(seedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// seedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// seedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	seedCmd.Flags().StringVar(&dbURL, "db_url", "", "postgres database url (required)")
	seedCmd.MarkFlagRequired("db_url")
}
