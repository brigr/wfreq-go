package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func containsString(s []string, e string) bool {
	// iterate elements in the slice and look for a match
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func do_wcount(inputdata string) (map[string]int64, []string, error) {
	var wcmap_keys []string

	// join adjacent lines separated by new line chars with white space
	inputdata = strings.ReplaceAll(inputdata, "\n", " ")

	// trim white space
	inputdata = strings.TrimSpace(inputdata)

	// split input data into space-delimited tokens
	tokens := strings.Split(inputdata, " ")

	// make a slice of all tokens and remove special chars
	for _, key := range tokens {
		// store token in no-repeat slice
		wcmap_keys = append(wcmap_keys, key)
	}

	// sort the slice
	sort.Slice(wcmap_keys, func(i, j int) bool {
		return strings.Compare(strings.ToLower(wcmap_keys[i]), strings.ToLower(wcmap_keys[j])) == -1
	})

	// iterate tokens and maintain word count
	wcmap := make(map[string]int64)

	for _, token := range wcmap_keys {
		// check if the token has been registered in the map
		_, ok := wcmap[token]

		if ok {
			// increase frequency count
			wcmap[token] += 1
		} else {
			// initialize frequency count
			wcmap[token] = 1
		}
	}

	return wcmap, wcmap_keys, nil
}

func do_print_freqs(wcmap map[string]int64, wcmap_keys_sorted []string) string {
	// keep the unique sorted keys
	var wcmap_keys_sorted_unique []string
	// keep the token keys
	var wcslice []string
	// calculate the output-line frame length to simulate Unix's tail
	var framelen int
	// output data in string format
	var output_data_str string

	for _, wcmap_key_sorted := range wcmap_keys_sorted {
		if containsString(wcmap_keys_sorted_unique, wcmap_key_sorted) {
			continue
		}

		wcmap_keys_sorted_unique = append(wcmap_keys_sorted_unique, wcmap_key_sorted)
	}

	// print count for each unique token from the string->int map
	for _, wcmap_key_sorted_unique := range wcmap_keys_sorted_unique {
		if wcmap_key_sorted_unique != "" {
			wcslice = append(wcslice, fmt.Sprintf("%7v %s", wcmap[wcmap_key_sorted_unique], wcmap_key_sorted_unique))
		}
	}

	// re-sort word count strings in the format "<wc> <w>"
	sort.Slice(wcslice, func(i, j int) bool {
		return strings.Compare(strings.ToLower(wcslice[i]), strings.ToLower(wcslice[j])) == -1
	})

	/*
	   Print the frame of the last 10 output lines if they exist;
	   otherwise print as many lines (less than 10) as possible.
	*/
	if len(wcslice) >= 10 {
		framelen = 9
	} else {
		framelen = len(wcslice) - 1
	}

	for i := len(wcslice) - 1 - framelen; i < len(wcslice); i++ {
		output_data_str += wcslice[i] + "\n"
	}

	return output_data_str
}

func main() {
	// set up logging
	log.SetPrefix("wfreq:")
	log.SetFlags(0)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/wordcount", func(c echo.Context) error {
		wcmap, wcmap_keys_sorted, _ := do_wcount(c.FormValue("text"))
		return c.HTML(http.StatusOK, do_print_freqs(wcmap, wcmap_keys_sorted))
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8083"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))

	// give up with success
	os.Exit(0)
}
