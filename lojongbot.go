package main

import (
	"encoding/json"
	"github.com/ChimeraCoder/anaconda"
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
	Sleep_time_in_hours         int
	Sleep_time_margin_in_hours  int
	Debug                       int
}

var logfile *os.File
var err error
var logger *log.Logger
var slogans []string
var num_slogans int
var configuration Configuration

// init opens a log file, reads the slogans file
// and creates a new random seed
func init() {

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

func tweet(tweet_text string) {
	// twitter api
	anaconda.SetConsumerKey(configuration.Twitter_consumer_key)
	anaconda.SetConsumerSecret(configuration.Twitter_consumer_secret)
	// I don't know about any possible timeout, therefore
	// initialize new for every tweet
	api := anaconda.NewTwitterApi(configuration.Twitter_access_token, configuration.Twitter_access_token_secret)

	if configuration.Debug == 1 {
		logger.Printf("DEBUG-MODE! I am not posting '%s'!", tweet_text)
	} else {
		tweet, err := api.PostTweet("#lojong slogan "+tweet_text, nil)
		if err != nil {
			logger.Printf("Problem posting '%s': %s", tweet_text, err)
		} else {
			logger.Printf("Tweet with slogan %s posted for user %s", tweet_text, tweet.User.ScreenName)
		}
	}
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

	// infinite loop
	for {

		tweet(slogans[rand.Intn(num_slogans)])
		sleep_hours := configuration.Sleep_time_in_hours + (configuration.Sleep_time_margin_in_hours - 2*rand.Intn(configuration.Sleep_time_margin_in_hours))
		logger.Printf("Will go to sleep for %d hours..", sleep_hours)
		time.Sleep(time.Duration(sleep_hours) * time.Hour)
	}
}
