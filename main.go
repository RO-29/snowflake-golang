package main

import (
	"fmt"
	"log"
	"time"

	"github.com/RO-29/snowflake-golang/snowflake"
)

func init() {
	snowflake.Init()
}

var sf = snowflake.NewSnowFlake()

//TODO add http server
func main() {

	for i := 0; i < 100; i++ {
		go generateUniqueSequence()
	}

	time.Sleep(100 * time.Second)
}

func generateUniqueSequence() {
	seqID, err := sf.GenerateUniqueSequenceID()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("seqID::", seqID)
}
