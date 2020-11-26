/*
Copyright © 2020 NAME HERE vinicius.costa.92@gmail.com

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
package http

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	address string
	watch   bool
	timeout int
	args    []string

	exitCode int
)

var Cmd = &cobra.Command{
	Use:   "http",
	Short: "http resource",
	Long:  "Checks wether a https endpoint is getting a valid response or not.",
	Run: func(cmd *cobra.Command, args []string) {
		httpCheck()
	},
}

func init() {
	// Cmd.Flags().StringVar(&address, "address", "localhost", "target postgres instance")
	Cmd.Flags().BoolVar(&watch, "watch", false, "keep watching command, retries connection each 2s.")
	Cmd.Flags().IntVar(&timeout, "timeout", 3600, "how many seconds should a watch run")

	args = Cmd.Flags().Args()
	if len(args) > 0 {
		address = args[0]
	} else {
		log.Println(Cmd.Long)
		os.Exit(1)
	}
}

func httpCheck() {

	c1 := make(chan int, 1)
	go func() {

		for {

			if response, err := http.Get(address); err != nil {
				log.Println(err)
				exitCode = 1
			} else {
				log.Printf("Connection sucessful!\n%d: %s", response.StatusCode, response.Status)
				exitCode = 0
				watch = false
			}

			if watch == false {
				break
			}
			time.Sleep(2 * time.Second)
		}
		c1 <- exitCode
	}()

	select {
	case res := <-c1:
		os.Exit(res)
	case <-time.After(time.Duration(timeout) * time.Second):
		log.Println("Timed out")
		os.Exit(127)
	}
}
