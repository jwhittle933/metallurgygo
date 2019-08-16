package main

import "time"

var (
	logger *Log
	start  time.Time
)

func init() {
	logger = StartLog()
	start = time.Now()
}

func main() {
	args := ParseFlags()

	NewFiles(args).
		Decode().
		Encode().
		Write()

	logger.I("Completed in %v", time.Now().Sub(start))
}
