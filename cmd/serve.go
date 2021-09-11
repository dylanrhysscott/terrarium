/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dylanrhysscott/terrarium/api/login"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the Terraform Registry",
	Long:  `Starts the Terraform Registry`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	loginAPI := login.NewLoginAPI("terraform-cli", "/oauth/auth", "/oauth/token", []int{10000})
	http.HandleFunc("/.well-known/terraform.json", loginAPI.DiscoveryHandler())
	port := ":8080"
	log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
