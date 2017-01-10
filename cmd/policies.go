// Copyright © 2017 G. Hussain Chinoy <ghchinoy@gmail.com>
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
	"github.com/spf13/cobra"
)

// policiesCmd represents the policies command
var policiesCmd = &cobra.Command{
	Use:   "policies",
	Short: "policies commands",
	Long: `commands to manage policies`,
}

func init() {
	RootCmd.AddCommand(policiesCmd)

	policiesCmd.PersistentFlags().StringVar(&cfgFile, "config", "", cfgHelp)
	policiesCmd.PersistentFlags().StringVar(&profile, "profile", "default", "profile name")
	policiesCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug output")
	// Set bash-completion
	validConfigFilenames := []string{"toml", ""}
	policiesCmd.PersistentFlags().SetAnnotation("config", cobra.BashCompFilenameExt, validConfigFilenames)


}
