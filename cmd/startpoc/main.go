package main

import (
	"fmt"
	"github.com/lcserny/go-authservice/src/logging"
	"time"
)

func main() {
	amount := 1000000

	startTime := time.Now()

	for i := range amount {
		logging.Info(fmt.Sprintf("iteration %d", i))
	}

	logging.Info(fmt.Sprintf("time taken for %d: %v", amount, time.Since(startTime)))
}
