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
	"time"

	"github.com/go-kit/kit/log"
	sql "github.com/honestbee/gopkg/sql/v2"
	"github.com/honestbee/gopkg/stringutils"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var (
	logger      log.Logger
	postgresDB  *sql.DB
	dbURL       string
	projectRoot = fmt.Sprintf("%s/src/github.com/honestbee/authorizer", os.Getenv("GOPATH"))

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "ctl",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		//	Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// initLogger inititialize logger
func initLogger() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller, "server", "ctl")
}

// initDB connects to test database
func initDB(dbURL string) *sql.DB {
	var err error
	if dbURL == "" {
		logger.Log("test", "database", "not", "found")
		os.Exit(1)
	}

	for count := 0; count < 5; count++ {
		d, err := sql.NewDB("postgres", dbURL)

		if err == nil {
			err = d.Ping()
		}
		if err != nil {
			logger.Log("postgres", "connect", "error", err, "retry", count < 5)
			time.Sleep(2 * time.Second)
			continue
		}
		// this converts all struct  name to snakecase name to map to db column
		d.MapperFunc(func(s string) string {
			return stringutils.ToLowerSnakeCase(s)
		})

		d.SetMaxIdleConns(2)
		d.SetMaxOpenConns(2)

		postgresDB = d
		return d
	}
	logger.Log("postgres", "connect", "error", err, "exiting", "...")
	os.Exit(1)
	return nil
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ctl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".ctl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ctl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
