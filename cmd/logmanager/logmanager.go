// Copyright (c) 2018 Zededa, Inc.
// All rights reserved.

package main

import (
	"bufio"
	"flag"
	"fmt"
	//"github.com/zededa/api/zmet"
	"github.com/zededa/go-provision/watch"
	"io"
	"log"
	"os"
	"strings"
)

const (
	logDirName = "/var/log"
)

var debug bool

// global stuff
var loggerReaderMap map[string]*bufio.Reader
type logReadHandler func(logFileName string, logContent string)
type logDeleteHandler func(logFileName string)

// Set from Makefile
var Version = "No version specified"

func main() {

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
	versionPtr := flag.Bool("v", false, "Version")
	debugPtr := flag.Bool("d", false, "Debug")
	flag.Parse()
	debug = *debugPtr
	if *versionPtr {
		fmt.Printf("%s: %s\n", os.Args[0], Version)
		return
	}

	log.Println("Starting log manager...")

	if loggerReaderMap == nil {
		log.Println("Creating loggerReader map")
		loggerReaderMap = make(map[string]*bufio.Reader)
	}

	logChanges := make(chan string)
	go watch.WatchStatus(logDirName, logChanges)
	log.Println("called watcher...")
	for {
		select {
		case change := <-logChanges:
			{
				//log.Println("change: ", change)
				HandleLogEvent(change, logDirName, handleLogModify, handleLogDelete)
			}
		}
	}
}

func HandleLogEvent(change string, logDirName string, handleLogModifyFunc logReadHandler, handleLogDeleteFunc logDeleteHandler) {

	operation := string(change[0])
	fileName := string(change[2:])
	if !strings.HasSuffix(fileName, ".log") {
		log.Printf("Ignoring file <%s> operation %s\n",
			fileName, operation)
		return
	}
	// Remove .log from name */
	name := strings.Split(fileName, ".log")
	logFileName := name[0]
	if operation == "D" {
		handleLogDeleteFunc(name[0])
		return
	}
	if operation != "M" {
		log.Fatal("Unknown operation from Watcher: ",
			operation)
	}
	logFilePath := logDirName + "/" + fileName
	go readLogFileLineByLine(logFilePath, logFileName, handleLogModifyFunc)

}
//XXX FIXME reader
func openLogFile(logFile string) *bufio.Reader {
	fileDesc, err := os.Open(logFile)

	if err != nil {
		log.Fatalf("%v for %s\n", err, logFile)
	}
	// Start reading from the file with a reader.
	reader := bufio.NewReader(fileDesc)
	if reader == nil {
		log.Fatalf("%s, reader create failed\n", logFile)
	}

	loggerReaderMap[logFile] = reader
	return reader
}

func getLoggerReader(logFile string) *bufio.Reader {
	reader, ok := loggerReaderMap[logFile]
	if !ok {
		return openLogFile(logFile)
	}
	return reader
}

//XXX FIXME reader.ReadString('\n') is not able to read any new changes
//made in the log files.
//we are checking that if we have an existing reader for any file(means we already 
//have opened the file) and if the file is opened then we are not opening it again
//for new change......this part is creating issue.
func readLogFileLineByLine(logFilePath,fileName string, handleLogModifyFunc logReadHandler) {

	logFile := logFilePath

	reader := getLoggerReader(logFile)
	if reader == nil {
		log.Fatalf("%s, log File open failed\n", logFile)
	}

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if debug {
				log.Println(err)
			}
			if err != io.EOF {
				fmt.Printf(" > Failed!: %v\n", err)
			}
			break
		}
		handleLogModifyFunc(fileName, line)
	}
}

func handleLogModify(logFilename string, logContent string) {
	//if debug {
		log.Printf("handleLogModify for %s\n", logFilename)
		log.Println("value of log content: ", logContent)
	//}
}

func handleLogDelete(logFilename string) {
	//if debug {
		log.Printf("handleLogDelete for %s\n", logFilename)
	//}

}
