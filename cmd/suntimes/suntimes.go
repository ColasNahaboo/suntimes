package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/nathan-osman/go-sunrise"
	// "regexp"
)

var usage = `Usage: %s latitude longitude [YYYY-MM-DD]

Prints the ISO times of sunrise and sunset, each on a line,
of the day YYYY-MM-DD, which defaults to today, at the given
latitude and longitude.

E.g: suntimes 43.6880520239146 -1.356146404748612 2022-02-22
     suntimes -- -12.2344 +23.12
     suntimes -S , 12.2344 23.12

Release: 1.0.0
Sources: https://github.com/ColasNahaboo/suntimes
Options:
`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
		flag.PrintDefaults()
	}
	spaceSep := flag.Bool("s", false, "Print sunrise and sunset times on the same line, space-separated")
	stringSep := flag.String("S", ``, "Print sunrise and sunset times on the same line, separated by the provided string")
	onlyRise := flag.Bool("sr", false, "Print only sunrise time")
	onlySet := flag.Bool("ss", false, "Print only sunset time")
	flag.Parse()
	var now time.Time
	var lat, lon float64
	var y, d int
	var m time.Month
	var err error
	var s []string
	switch flag.NArg() {
	case 3:
		reday := regexp.MustCompile("([[:digit:]]{4})-([[:digit:]]{2})-([[:digit:]]{2})")
		if s = reday.FindStringSubmatch(flag.Arg(2)); s == nil {
			log.Fatalf(`Arg#3 "%v" is not a day (must be YYYY-MM-DD)`, flag.Arg(2))
		}
		y, err = strconv.Atoi(s[1])
		if err != nil {
			log.Fatalf(`Arg#3 "%v" year parse error, %v`, flag.Arg(2), err)
		}
		mn, err := strconv.Atoi(s[2])
		if err != nil {
			log.Fatalf(`Arg#3 "%v" month parse error, %v`, flag.Arg(2), err)
		}
		m = time.Month(mn)
		d, err = strconv.Atoi(s[3])
		if err != nil {
			log.Fatalf(`Arg#3 "%v" day parse error, %v`, flag.Arg(2), err)
		}
		fallthrough
	case 2:
		lat, err = strconv.ParseFloat(flag.Arg(0), 64)
		if err != nil {
			log.Fatalf(`Arg#1 "%v" (latitude) must be a floating point number, %v`, flag.Arg(0), err)
		}
		lon, err = strconv.ParseFloat(flag.Arg(1), 64)
		if err != nil {
			log.Fatalf(`Arg#2 "%v" (longitude) must be a floating point number, %v`, flag.Arg(1), err)
		}
	default:
		log.Fatalf(`Bad number of arguments: %v. Use -h for help.`, flag.NArg())
	}
	if y == 0 {
		now = time.Now()
		y, m, d = now.Date()
	}
	rise, set := sunrise.SunriseSunset(lat, lon, y, m, d)

	if *onlyRise {
		fmt.Println(rise.Format(time.RFC3339))
	} else if *onlySet {
		fmt.Println(set.Format(time.RFC3339))
	} else if *spaceSep {
		fmt.Println(rise.Format(time.RFC3339), set.Format(time.RFC3339))
	} else if *stringSep != `` {
		fmt.Printf("%v%v%v\n", rise.Format(time.RFC3339), *stringSep, set.Format(time.RFC3339))
	} else {
		fmt.Println(rise.Format(time.RFC3339))
		fmt.Println(set.Format(time.RFC3339))
	}
}
