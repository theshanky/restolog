package main

import (
	"fmt"
	"os"
	cmd "restolog/resto"
)

func main() {
	// Open the log file for reading
	file, err := os.Open("testdata/log.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	count, err := cmd.RestoLog(file)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("Top 3 menu items: ", count)
}
