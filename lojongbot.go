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

var logfile *os.File
var err error
var logger *log.Logger

func init_bot() {
	// logging
	logfile, err = os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	check(err)
	logger = log.New(logfile, "", log.Ldate|log.Ltime|log.Lshortfile)
	rand.Seed(time.Now().UnixNano())

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	init_bot()

	content, err := ioutil.ReadFile(sourceFileName)
	check(err)
	//fmt.Print(string(dat))
	lines := strings.Split(string(content), "\n")

	num_lines := len(lines) - 1
	logger.Printf("Found %d elements in %s\n", num_lines, sourceFileName)

	fmt.Printf("%s\n", lines[rand.Intn(num_lines)])
}
