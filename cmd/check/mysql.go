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
	"log"
	"os"
	"time"
	"database/sql"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

var mysqlInstructions = `Check if a mysql instance is responding correctly.
`

var MYSQLCmd = &cobra.Command{
	Short:   "check mysql databases",
	Example: "endpointer check mysql --port 3306",
	Use:     "mysql <url>",
	Args:    cobra.ExactArgs(1),
	Long:    mysqlInstructions,
	Run: func(cmd *cobra.Command, args []string) {
		mysqlCheck(args)
	},
}

func init() {
	MYSQLCmd.Flags().StringVar(&port, "port", "3306", "target mysql port")
	MYSQLCmd.Flags().StringVar(&user, "user", "root", "target mysql username")
	MYSQLCmd.Flags().StringVar(&password, "password", "password", "target users password")
	MYSQLCmd.Flags().StringVar(&database, "database", "database", "target database")
	MYSQLCmd.Flags().BoolVar(&watch, "watch", false, "keep watching command, retries connection each 2s.")
	MYSQLCmd.Flags().IntVar(&timeout, "timeout", 3600, "how many seconds should a watch run")

}

func mysqlCheck(args []string) {

	if len(args) > 0 {
		address = args[0]

	} else {
		log.Println(mysqlInstructions)
		os.Exit(1)
	}

	c1 := make(chan int, 1)

	go func() {

		for{
		db, err := sql.Open("mysql", user + ":" + password + "@tcp(" + address + ":" + port + ")/"  + database)

		if err != nil {
			log.Println(err)
	
		}	

		err = db.Ping()	

		if err != nil {
			log.Println("Connection failed!")
			exitCode = 1

		}else {
			log.Println("Connection sucessful!")
			exitCode = 0
			watch = false
		}

		defer db.Close()

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