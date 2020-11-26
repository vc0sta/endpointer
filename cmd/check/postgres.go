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
package check

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/spf13/cobra"
)

var (
	address  string
	port     string
	user     string
	password string
	database string
	watch    bool
	timeout  int

	exitCode int
)

var PostgresCmd = &cobra.Command{
	Use:   "postgres",
	Short: "postgres resource",
	Long:  "Prints stuff about the user. You could also use the flags in your addPartner() function",
	Run: func(cmd *cobra.Command, args []string) {
		postgresCheck()
	},
}

func init() {
	PostgresCmd.Flags().StringVar(&address, "address", "localhost", "target postgres instance")
	PostgresCmd.Flags().StringVar(&port, "port", "5432", "target postgres port")
	PostgresCmd.Flags().StringVar(&user, "user", "postgres", "target postgres user")
	PostgresCmd.Flags().StringVar(&password, "password", "postgres", "target users password")
	PostgresCmd.Flags().StringVar(&database, "database", "postgres", "target database")
	PostgresCmd.Flags().BoolVar(&watch, "watch", false, "keep watching command, retries connection each 2s.")
	PostgresCmd.Flags().IntVar(&timeout, "timeout", 3600, "how many seconds should a watch run")
}

func postgresCheck() {

	c1 := make(chan int, 1)
	go func() {

		for {
			pgdb := pg.Connect(&pg.Options{
				Addr:     address + ":" + port,
				User:     user,
				Password: password,
				Database: database,
			})

			ctx := context.Background()

			if err := pgdb.Ping(ctx); err != nil {
				log.Println(err)
				exitCode = 1
			} else {
				log.Println("Connection sucessful!")
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
