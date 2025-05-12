package main

import (
    "strings"
    "errors"
    "sort"
    "log"
    "fmt"
    "os"
    "io"
)

type frame struct {
    wcount int64
    word string
}

func containsString(s []string, e string) bool {
    // iterate elements in the slice and look for a match
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func readDataFromKeyboard() (string, bool) {
    // data buffer
    var alltext = ""
    // keep error
    var err error = nil
    
    // allocate 1 MB per line
    buf := make([]byte, 1024 * 1024 * 1024)

    for {
	// read data from stdin
        n, err := os.Stdin.Read(buf)

	// append current input data to buffer
        alltext += string(buf[:n])

	// expect EOF to end I/O
	if err == io.EOF {
	    break
	}

    }

    return alltext, err == io.EOF
}

func do_read_tokens() (string, error) {
    // read input string
    itxt, err := readDataFromKeyboard()
    if err {
        return "", errors.New("I/O error: No input text was given")
    }
    
    // join adjacent lines separated by new line chars with white space
    itxt = strings.ReplaceAll(itxt, "\n", " ")

    // trim white space
    itxt = strings.TrimSpace(itxt)

    return itxt, nil
}

func do_wcount(itxt string, ignore_special_chars bool) (map[string]int64,[]string,error) {
    var wcmap_keys []string
    
    // split input data into space-delimited tokens
    tokens := strings.Split(itxt, " ")

    // make a slice of all tokens and remove special chars
    for _, key := range tokens {
            if ignore_special_chars {
	        key = strings.ReplaceAll(key, ".", "")
	        key = strings.ReplaceAll(key, ",", "")
	        key = strings.ReplaceAll(key, ";", "")
	        key = strings.ReplaceAll(key, "\"", "")
            }
            
	    // store token in no-repeat slice
	    wcmap_keys = append(wcmap_keys, key)
    }

    // sort the slice
    sort.Slice(wcmap_keys, func(i, j int) bool { return strings.Compare(strings.ToLower(wcmap_keys[i]),strings.ToLower(wcmap_keys[j])) == -1 })

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

    return wcmap,wcmap_keys,nil
}

func do_print_freqs(wcmap map[string]int64, wcmap_keys_sorted []string, allow_stdout bool) []frame {
    // keep the unique sorted keys
    var wcmap_keys_sorted_unique []string
    // keep the token keys
    var wcslice []string
    // calculate the output-line frame length to simulate Unix's tail
    var framelen int
    // var output frames
    var output_frame []frame
    // var tail'ed output frames
    var output_frame_tail []frame

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
            output_frame = append(output_frame, frame{wcount: wcmap[wcmap_key_sorted_unique], word: wcmap_key_sorted_unique})
        }
    }

    // re-sort word count strings in the format "<wc> <w>"
    sort.Slice(wcslice, func(i, j int) bool { return strings.Compare(strings.ToLower(wcslice[i]),strings.ToLower(wcslice[j])) == -1 })

    /* 
       Print the frame of the last 10 output lines if they exist;
       otherwise print as many lines (less than 10) as possible.
    */
    if len(wcslice) >= 10 {
	framelen = 9
    } else {
        framelen = len(wcslice) - 1
    }

    for i := len(wcslice)-framelen-1; i < len(wcslice); i++ {
        if allow_stdout {
            fmt.Println(wcslice[i])
        }
        
        output_frame_tail = append(output_frame_tail, output_frame[i])
    }

    return output_frame_tail
}


func main() {
    var ignore_special_chars = false
    var allow_stdout = true

    // set up logging
    log.SetPrefix("wfreq:")
    log.SetFlags(0)

    // read EOF-ended string from keyboard
    itxt, err := do_read_tokens()

    if err != nil {
        log.Println("Error while reading data from keyboard")
        os.Exit(1)
    }

    // run word count loop
    wcmap, wcmap_keys_sorted, err := do_wcount(itxt, ignore_special_chars)
    
    if err != nil {
        log.Println("Error while expecting data from keyboard")
	os.Exit(1)
    }
    
    // make report to stdout
    _ = do_print_freqs(wcmap, wcmap_keys_sorted, allow_stdout)

    // give up with success
    os.Exit(0)
}
