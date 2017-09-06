// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"log"

	"github.com/gotoolkit/peony/http"
	"github.com/gotoolkit/peony/jwt"
	"github.com/gotoolkit/peony/store"

	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start api",

	Run: func(cmd *cobra.Command, args []string) {
		db := store.NewStore(
			store.Dialect("mysql"),
			store.Args(fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, passwd, dbname)),
		)

		err := db.Open()
		if err != nil {
			log.Fatalf("Fail to init db: %v", err)
		}
		defer db.Close()

		jwtService, err := jwt.NewService()
		if err != nil {
			log.Fatalf("Fail to init jwt service: %v", err)
		}

		server := &http.Server{
			BindAddress: addr,
			UserService: db.UserService,
			Debug:       debug,
			JWTService:  jwtService,
		}
		err = server.Start()
		if err != nil {
			log.Fatalf("Fail to start server: %v", err)
		}
	},
}

var (
	username string
	passwd   string
	dbname   string
	addr     string
	debug    bool
)

func init() {
	RootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	apiCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Debug mode")
	apiCmd.Flags().StringVarP(&addr, "addr", "a", ":8080", "Bind address for api")
	apiCmd.Flags().StringVarP(&username, "user", "u", "root", "User name for database")
	apiCmd.Flags().StringVarP(&passwd, "pass", "p", "root", "Password for database")
	apiCmd.Flags().StringVarP(&dbname, "name", "n", "peony", "database name for database")
}
