// package main

// import (
// 	"fmt"
// 	"github.com/jrallison/go-workers"
// )

// func myJob(message *workers.Msg) {
// 	// do something with your message
// 	// message.Jid()
// 	// message.Args() is a wrapper around go-simplejson (http://godoc.org/github.com/bitly/go-simplejson)
// 	fmt.Println(message)
// }

// type myMiddleware struct{}

// func (r *myMiddleware) Call(queue string, message *workers.Msg, next func() bool) (acknowledge bool) {
// 	// do something before each message is processed
// 	acknowledge = next()
// 	// do something after each message is processed
// 	return
// }

// func main() {
// 	workers.Configure(map[string]string{
// 		// location of redis instance
// 		"server": "localhost:6379",
// 		// instance of the database
// 		"database": "0",
// 		// number of connections to keep open with redis
// 		"pool": "30",
// 		// unique process id for this instance of workers (for proper recovery of inprogress jobs on crash)
// 		"process": "1",
// 	})

// 	workers.Middleware.Append(&myMiddleware{})

// 	// pull messages from "myqueue" with concurrency of 10
// 	workers.Process("myqueue", myJob, 10)

// 	// pull messages from "myqueue2" with concurrency of 20
// 	workers.Process("myqueue2", myJob, 20)

// 	// Add a job to a queue
// 	workers.Enqueue("myqueue3", "Add", []int{1, 2})

// 	// Add a job to a queue with retry
// 	workers.EnqueueWithOptions("myqueue3", "Add", []int{1, 2}, workers.EnqueueOptions{Retry: true})

// 	// stats will be available at http://localhost:8080/stats
// 	go workers.StatsServer(8080)

// 	// Blocks until process is told to exit via unix signal
// 	workers.Run()
// }
package main

import (
	"fmt"
	"github.com/jrallison/go-workers"
	// "log"
	"models"
)
import "github.com/zebresel-com/mongodm"

// var config = &bongo.Config{
// 	ConnectionString: "localhost",
// 	Database:         "super_control_development",
// }

// var connection = bongo.Connect(&config)
var dbConfig = &mongodm.Config{
	DatabaseHosts:    []string{"127.0.0.1"},
	DatabaseName:     "super_control_development",
	DatabaseUser:     "owner",
	DatabasePassword: "1qazxsw2",
	// The option `DatabaseSource` is the database used to establish
	// credentials and privileges with a MongoDB server. Defaults to the value
	// of `DatabaseName`, if that is set, or "admin" otherwise.
	DatabaseSource: "super_control_development",
}

func GoWorker(message *workers.Msg) {
	fmt.Println(message)
	connection, err := mongodm.Connect(dbConfig)
	defer connection.Close()
	if err != nil {
		fmt.Println("Database connection error: %v", err)
	}
	connection.Register(&models.Person{}, "persons")

	Person := connection.Model("Person")

	person := &models.Person{}

	Person.New(person) //this sets the connection/collection for this type and is strongly necessary(!) (otherwise panic)

	person.FirstName = "Tien An"
	person.LastName = "Vo"
	person.Gender = "male"
	person.Save()
}

func main() {
	workers.Configure(map[string]string{
		// location of redis instance
		"server": "localhost:6379",
		// instance of the database
		"database": "0",
		// number of connections to keep open with redis
		"pool": "10",
		// unique process id for this instance of workers (for proper recovery of inprogress jobs on crash)
		"process": "1",
	})

	// pull messages from "go_queue" with concurrency of 10
	workers.Process("go_queue", GoWorker, 10)

	// Blocks until process is told to exit via unix signal
	workers.Run()
}
