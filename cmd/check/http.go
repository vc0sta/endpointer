/*Copyright © 2020 NAME HERE vinicius.costa.92@gmail.com

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
package check

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var instructions = `Check if a http/https is reachable.
It will not test if the returning code is 2XX.`

var HttpCmd = &cobra.Command{
	Use:     "http <url>",
	Short:   "check http/https endpoints",
	Args:    cobra.ExactArgs(1),
	Long:    instructions,
	Example: "endpointer check http https://google.com",
	Run: func(cmd *cobra.Command, args []string) {
		httpCheck(args)
	},
}

func init() {
	HttpCmd.Flags().BoolVar(&watch, "watch", false, "keep watching command, retries connection each 2s.")
	HttpCmd.Flags().IntVar(&timeout, "timeout", 3600, "how many seconds should a watch run")

}

func httpCheck(args []string) {

	if len(args) > 0 {
		address = args[0]
	} else {
		log.Println(instructions)
		os.Exit(1)
	}

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
