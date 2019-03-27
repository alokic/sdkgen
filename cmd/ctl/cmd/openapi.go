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
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

type oaOpt struct {
	in  string
	out string
}

var (
	oaOpts = &oaOpt{}

	// openapiCmd represents the openapi generate command
	openapiCmd = &cobra.Command{
		Use:   "openapi",
		Short: "Generate openapi specs from jsons",
		Long: `Generate openapi 3 spec from json request and responses of REST APIs. For example:

	ctl openapi --in examples/input --out examples/out`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := checkDirectoryExist(oaOpts.in); err != nil {
				fmt.Printf("openapi generation error: input folder not found: %v\n", err)
				os.Exit(1)
			}

			if err := checkDirectoryExist(oaOpts.out); err != nil {
				fmt.Printf("openapi generation error: output folder not found: %v\n", err)
				os.Exit(1)
			}

			commmand := exec.Command("sdkgen", "--in", oaOpts.in, "--out", oaOpts.out)
			err := commmand.Run()

			var out bytes.Buffer
			commmand.Stdout = &out

			var errb bytes.Buffer
			commmand.Stderr = &errb

			if err != nil {
				fmt.Printf("openapi generation error: %v %v\n", err.Error(), err.(*exec.ExitError).Stderr)
				fmt.Println(errb.String())
				os.Exit(1)
			}
			fmt.Println("openapi generation success")
			fmt.Println(out.String())
		},
	}
)

func init() {
	rootCmd.AddCommand(openapiCmd)

	openapiCmd.Flags().StringVarP(&oaOpts.in, "in", "i", defaultInputPath, "path where json specs are kept (required).")
	openapiCmd.Flags().StringVarP(&oaOpts.out, "out", "o", defaultOutputPath, "output path (required).")
}
