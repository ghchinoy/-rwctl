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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"github.com/ghchinoy/rwctl/portal"
	"github.com/ghchinoy/rwctl/control"
)

// portalresetCmd represents the portalreset command
var portalresetCmd = &cobra.Command{
	Use:   "reset [theme]",
	Short: "reset portal theme",
	Long: `resets portal theme to vanilla.
The theme can be specified on the command-line as an argument or via the profile's configuration variable.`,
	Run: func(cmd *cobra.Command, args []string) {
		// obtain profile values
		profilemap := viper.GetStringMap(profile)
		// validate that 'theme' value exists
		_, ok := profilemap["theme"]

		// set theme
		var theme string
		if len(args) == 0 { // no theme specified, use profile
			if ok {
				theme = profilemap["theme"].(string)
			} else { // no profile theme set
				theme = "hermosa"
			}
		} else { // use theme from args
			theme = args[0]
		}

		// change profilemap into control.Configuration
		config, err := control.ViperToConfiguration(profilemap, debug)
		if err != nil {
			fmt.Println("Error converting configuration", err.Error())
			os.Exit(1)
		}

		// execute rebuild styles
		portal.ResetTheme(config, theme, debug)
	},
}

func init() {
	portalCmd.AddCommand(portalresetCmd)



}
