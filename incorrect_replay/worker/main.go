package main

import (
	"log"
	"temporal_replay/incorrect_replay"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "hello-world", worker.Options{})

	w.RegisterWorkflow(incorrect_replay.TestWorkflow)
	w.RegisterActivity(incorrect_replay.ActivityA)
	w.RegisterActivity(incorrect_replay.ActivityB)
	w.RegisterActivity(incorrect_replay.ActivityC)
	w.RegisterActivity(incorrect_replay.ActivityD)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
