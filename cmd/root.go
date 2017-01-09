// Copyright © 2016 G. Hussain Chinoy <ghchinoy@gmail.com>
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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
const cfgHelp = `config file (default is $HOME/.config/roguewave/rwctl.toml)`
var profile string
var VERSION string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "rwctl",
	Short: "Rogue Wave API Platform command-line interface control",
	Long: `rwctl is a CLI tool to manage the Rogue Wave API Platform, including
 APIs, Policies, and API platform and portal settings.`,

// Uncomment the following line if your bare application
// has an action associated with it:
//	Run: func(cmd *cobra.Command, args []string) { },
}

// SetVersion sets the version for use within cmd package
func SetVersion(v string) {
	VERSION = v
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	//RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rwctl.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if profile != "" {
		viper.Set("profile", profile)
	}

	viper.SetConfigName("rwctl") // name of config file (without extension)
	viper.AddConfigPath("$HOME")  // adding home directory as second search path
	viper.AddConfigPath("$HOME/.config/roguewave")  // adding .config directory's roguewave as first search path
	viper.AddConfigPath(".") // local
	viper.AutomaticEnv()          // read in environment variables that match

	if cfgFile != "" { // enable ability to specify config file via flag
		fmt.Println("non blank cfg:", cfgFile)
		viper.SetConfigFile(cfgFile)
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())

		// shouldn't need to set cfgFile globally; once viper reads it (this block)
		// all the values are accessible via the global viper struct
		cfgFile = viper.ConfigFileUsed()
	} else {
		fmt.Println("Could not find a", cfgHelp)
		os.Exit(1)
	}
}
