package measure

import (
	"log"
	"time"
)

func Time(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
