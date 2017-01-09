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
	"os"
	"github.com/spf13/viper"
	"github.com/ghchinoy/rwctl/portal"
	"github.com/ghchinoy/rwctl/control"
)



// uploadCmd uploads a file or files to the Portal's CMS path
var uploadCmd = &cobra.Command{
	Use:   "upload <file...>",
	Short: "upload a file or files to the portal cms",
	Long: `upload a file or files to the portal cms`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("a file (or files) argument is required. See -h help.")
			os.Exit(1)
		}
		if cmspathtarget == "" {
			fmt.Println("a CMS path target defined by --path flag is required. See -h help.")
			os.Exit(1)
		}
		//fmt.Println("uploading", args, "to", cmspathtarget)

		var cfgmap map[string]interface{}
		var config control.Configuration

		if viper.IsSet(profile) {
			cfgmap = viper.GetStringMap(profile)
		} else {
			fmt.Println("Cannot find profile", profile, "in configuration.")
			os.Exit(1)
		}

		config, err := control.ViperToConfiguration(cfgmap, debug)
		if err != nil {
			fmt.Println("Error translating config", err.Error())
		} else {
			portal.UploadFiles(args, config, cmspathtarget, debug)
		}
	},
}

func init() {
	portalCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	uploadCmd.Flags().StringVarP(&cmspathtarget, "path", "p" , "","CMS path target")

}
