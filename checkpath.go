package main

import (
	"os"
	"io"
	"log"
	"flag"
)

func main() {
	osPath := flag.String("p", "", "Path or file to check")
	flag.Parse()

	logFile, err := os.OpenFile("checkpath.log", os.O_CREATE | os.O_APPEND | os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	if _, err := os.Stat(*osPath); os.IsNotExist(err){
		log.Println("ERROR: The path does not exist.", err)
	}

	if _, err := os.Stat(*osPath); err == nil {
		log.Println("The path exists: ", *osPath)
	}
}
