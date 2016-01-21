package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type Configuration struct {
        Sourcefile                  string
	Logfile                     string
	Twitter_access_token        string
	Twitter_access_token_secret string
	Twitter_consumer_key        string
	Twitter_consumer_secret     string
        Sleep_time_in_seconds int
}


var logfile *os.File
var err error
var logger *log.Logger
var slogans []string
var num_slogans int
var configuration Configuration

// init_bot opens a log file, reads the slogans file
// and creates a new random seed
func init_bot() {

	// config
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration = Configuration{}
	err := decoder.Decode(&configuration)
	check(err)

	// logging
	logfile, err = os.OpenFile(configuration.Logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	check(err)
	logger = log.New(logfile, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Print("Started...")

	// init random generator
	rand.Seed(time.Now().UnixNano())

	// read source file
	content, err := ioutil.ReadFile(configuration.Sourcefile)
	check(err)
	//fmt.Print(string(dat))
	slogans = strings.Split(string(content), "\n")

	num_slogans = len(slogans) - 1
	logger.Printf("Found %d elements in %s\n", num_slogans, configuration.Sourcefile)

}

// check panics if an error is detected
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// catch interrupts
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		logger.Print("ended...")
		os.Exit(1)
	}()

	init_bot()

	for i := 0; i < 10; i++ {

		fmt.Printf("%s\n", slogans[rand.Intn(num_slogans)])
		logger.Printf("Will go to sleep for %d seconds..", configuration.Sleep_time_in_seconds)
		time.Sleep(time.Duration(configuration.Sleep_time_in_seconds) * time.Second)
	}
}
