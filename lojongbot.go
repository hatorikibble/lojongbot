package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const sourceFileName string = "slogans.txt"
const logFileName string = "lojongbot.log"
const sleepTime int = 10

var logfile *os.File
var err error
var logger *log.Logger
var slogans []string
var num_slogans int

// init_bot opens a log file, reads the slogans file
// and creates a new random seed
func init_bot() {
	// logging
	logfile, err = os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	check(err)
	logger = log.New(logfile, "", log.Ldate|log.Ltime|log.Lshortfile)

	// init random generator
	rand.Seed(time.Now().UnixNano())

	// read source file
	content, err := ioutil.ReadFile(sourceFileName)
	check(err)
	//fmt.Print(string(dat))
	slogans = strings.Split(string(content), "\n")

	num_slogans = len(slogans) - 1
	logger.Printf("Found %d elements in %s\n", num_slogans, sourceFileName)

}

// check panics if an error is detected
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	init_bot()

	for i := 0; i < 10; i++ {

		fmt.Printf("%s\n", slogans[rand.Intn(num_slogans)])
		logger.Printf("Will go to sleep for %d seconds..", sleepTime)
		time.Sleep(time.Duration(sleepTime) * time.Second)
	}
}
