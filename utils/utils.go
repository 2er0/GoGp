package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"
)

// PseudoUuid returns a new uuid
func PseudoUuid() (uuid string) {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return
}

func BaseLogInfo() int {
	return log.Ldate | log.Ltime
}

// TimeTrack returns elapsed time and puts on the given channel
func TimeTrack(start time.Time, name string, gens int, st chan interface{}) time.Duration {
	elapsed := time.Since(start)
	st <- fmt.Sprintf("%s INFO | Execution of Generation %d took %s", name, gens, elapsed)
	return elapsed
}
