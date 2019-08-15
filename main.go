package main

func main() {
	args := ParseFlags()
	NewFiles(args).
		Decode().
		Encode().
		Write()
}
