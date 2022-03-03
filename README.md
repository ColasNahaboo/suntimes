# suntimes
A small command-line utility to display sunrise and sunset times, useful for scripts in bash and other shells.

It is just an ultra-simple wrapper around https://github.com/nathan-osman/go-sunrise

## Usage
`suntimes latitude longitude [YYYY-MM-DD]`

Prints the ISO times of sunrise and sunset, each on a line,
of the day YYYY-MM-DD, which defaults to today, at the given
latitude and longitude.

Options:
- `-S string`
        Print sunrise and sunset times on the same line, separated by the provided string
- `-s`
        Print sunrise and sunset times on the same line, space-separated
- `-sr`
        Print only sunrise time
-  `-ss`
        Print only sunset time

Examples:
```
$ suntimes 43.6880520239146 -1.356146404748612 2022-02-22
2022-02-22T06:54:01Z
2022-02-22T17:44:12Z

$ suntimes -- -12.2344 +23.12
2022-03-03T04:30:16Z
2022-03-03T16:49:05Z

$ suntimes -s 12.34 56.78
2022-03-03T02:27:42Z 2022-03-03T14:22:24Z

$ suntimes -S , 12.34 56.78
2022-03-03T02:27:42Z,2022-03-03T14:22:24Z

$ suntimes -sr 12.34 56.78
2022-03-03T02:27:42Z
```

## Installation
- Compile it byt going into `cmd/suntimes` and typing `go build`
  You may have to run also a `go get github.com/nathan-osman/go-sunrise` if prompted.
- copy the generated executable in your PATH, e.g. `cp suntimes /usr/local/bin`

## History
- v1.0.0 2022-03-03 Initila release


(c) 2022 Colas Nahaboo, MIT license.
