package main

import "fmt"

func main() {
	fmt.Println("============ Starting Meme API ============");

	loadConfig()
	setupMysql()
	apiSetup()

}
