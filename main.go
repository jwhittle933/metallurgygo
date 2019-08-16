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

	// Single chain of operations
	// Error checking is done at every step
	// If error, log error and os.Exit
	NewFiles(args).Decode().Encode().Write()

	logger.I("Completed in %v", time.Now().Sub(start))
}

