package main

import (
	"errors"
	"fmt"
	"gomp"
	"log"
	"math/rand"
)

const jobId = "asfsdgsdg"
const otherJobId = "gfsdgds"

func main() {
	serv, err := gomp.NewService("localhost:61613", nil)
	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < 10; i++ {
		err := serv.SendMessage(jobId, []byte(fmt.Sprintf("%d", i)))
		if err != nil {
			log.Println(err)
			return
		}
	}

	log.Println("message sent successfully, now moving forward to ingestion")

	serv.RegisterFuncWithJobID(jobId, func(m []byte) error {
		fmt.Println(string(m))
		if rand.Int()%2 == 0 {
			return errors.New("new random error passed")
		}
		return nil
	})

	if err := serv.StartMultiFuncProcessing(); err != nil {
		log.Println(err)
		return
	}
}
