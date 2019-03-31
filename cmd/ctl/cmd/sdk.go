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
	"strings"

	"github.com/spf13/cobra"
)

type sdkOpt struct {
	in     string
	out    string
	lang   string
	config string
}

var (
	sdkOpts = &sdkOpt{}

	// sdkCmd represents the openapi generate command
	sdkCmd = &cobra.Command{
		Use:   "sdk",
		Short: "Generate sdk from openapi specs.",
		Long: `Generate sdk from opeaapi specs. 
		Its a wrapper around https://github.com/OpenAPITools/openapi-generator.
		For example:

	ctl sdk -i examples/out/main.yaml -o examples/sdk/python -g python -c ./lang/python.json
		
		If you have installed openapi-generator from https://github.com/OpenAPITools/openapi-generator, you can run as below to generate sdk:
	
	openapi-generator generate   -i examples/out/main.yaml -o examples/sdk/python -g python -c  ./lang/python.json --enable-post-process-file 
	`,
		Run: func(cmd *cobra.Command, args []string) {

			if err := checkDirectoryExist(sdkOpts.in); err != nil {
				fmt.Printf("sdk generation error: input folder not found: %v\n", err)
				os.Exit(1)
			}

			if err := checkDirectoryExist(sdkOpts.out); err != nil {
				fmt.Printf("sdk generation error: output folder not found: %v\n", err)
				os.Exit(1)
			}

			os.RemoveAll(sdkOpts.out)

			opts := []string{"generate", "-i", sdkOpts.in, "-o", sdkOpts.sdkOutputPath(), "-g", sdkOpts.lang, "--enable-post-process-file"}

			if sdkOpts.config != "" {
				opts = append(opts, "-c", sdkOpts.config)
			}

			var commmand *exec.Cmd
			if os.Getenv("ENV_TYPE") == "docker" {
				commmand = exec.Command("docker-entrypoint.sh", opts...)
			} else {
				commmand = exec.Command("openapi-generator", opts...)
			}

			var out bytes.Buffer
			commmand.Stdout = &out

			var errb bytes.Buffer
			commmand.Stderr = &errb

			err := commmand.Run()
			if err != nil {
				fmt.Printf("openapi generation error: %v\n", err)
				fmt.Println(errb.String())
				os.Exit(1)
			}
			fmt.Println("openapi generation success")
			fmt.Println(out.String())
		},
	}
)

func (s *sdkOpt) sdkOutputPath() string {
	return strings.TrimRight(sdkOpts.out, "/") + "/" + sdkOpts.lang
}

func init() {
	rootCmd.AddCommand(sdkCmd)

	sdkCmd.Flags().StringVarP(&sdkOpts.in, "in", "i", defaultInputPath, "path of openapi YAML file. (required).")
	sdkCmd.MarkFlagRequired("in")

	sdkCmd.Flags().StringVarP(&sdkOpts.out, "out", "o", defaultOutputPath, "output path where sdk will be generated(required).")
	sdkCmd.MarkFlagRequired("out")

	sdkCmd.Flags().StringVarP(&sdkOpts.lang, "lang", "g", "", "language in which sdk should be generated(required).")
	sdkCmd.MarkFlagRequired("lang")

	sdkCmd.Flags().StringVarP(&sdkOpts.config, "conf", "c", "", `language specific config for sdk generation. 
		check https://github.com/OpenAPITools/openapi-generator for details`)
}
