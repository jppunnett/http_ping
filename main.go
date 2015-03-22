package main

// Command http_ping pings a Web site. Use it to test if a Web site is up.
//
// See http_ping -help for instructions

import (
	"fmt"
	"flag"
	"os"
	"net/http"
	"time"
	"io/ioutil"
)

var (
	outputResp = flag.Bool("print", false, "Output page response")
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %v [flags] <URL>\n", os.Args[0])
		flag.PrintDefaults()
	}
	
	flag.Parse()
	
	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(2)
	}

	url := flag.Arg(0)	

	fmt.Fprintf(os.Stdout, "Pinging %v...\n", url)

	t0 := time.Now()
	resp, err := http.Get(url)
	t1 := time.Now()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error pinging %v: %v\n", url, err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	fmt.Fprintf(os.Stdout, "Got response from %v in %v.\n", url, t1.Sub(t0))
	if *outputResp {
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			fmt.Fprintf(os.Stdout, "%s", body)
		}
	}
}