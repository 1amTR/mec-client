package main

import (
    "log"
    "time"
)

func runningtime(s string) (string, time.Time) {
    log.Println("Start:	", s)
    return s, time.Now()
}

func track(s string, startTime time.Time) {
    endTime := time.Now()
    log.Println("End:	", s, "took", endTime.Sub(startTime))
}

func main() {
	runningtime("execute")	
	track()
}
