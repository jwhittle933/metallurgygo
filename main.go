package main

var logger *Log

func init() {
	logger = StartLog()
}

func main() {
	args := ParseFlags()
	NewFiles(args).
		Decode().
		Encode().
		Write()
}
