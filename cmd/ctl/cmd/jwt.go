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
	"strings"

	"github.com/honestbee/gopkg/auth"
	"github.com/spf13/cobra"
)

var (
	jwtClaims string
	jwtKey    string
	// jwtCmd represents the jwt command
	jwtCmd = &cobra.Command{
		Use:   "jwt",
		Short: "Generate a jwt token",
		Long: `Generate a jwt token. For example:

	ctl jwt --claims 'user_id=1, email=a@b.com'`,
		Run: func(cmd *cobra.Command, args []string) {
			claims := map[string]interface{}{}
			for _, v0 := range strings.Split(jwtClaims, ",") {
				e := strings.Split(strings.TrimSpace(v0), "=")
				claims[e[0]] = e[1]
			}
			j := auth.NewJWT(jwtKey)
			token, err := j.Generate(claims)
			if err != nil {
				logger.Log("error", err)
				os.Exit(1)
			}
			fmt.Println(token)
		},
	}
)

func init() {
	rootCmd.AddCommand(jwtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jwtCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jwtCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	jwtCmd.Flags().StringVar(&jwtKey, "key", "", "jwt key (required).")
	jwtCmd.MarkFlagRequired("key")

	jwtCmd.Flags().StringVar(&jwtClaims, "claims", "", "jwt claims (required).")
	jwtCmd.MarkFlagRequired("claims")
}
