package main

import (
    "testing"
)

func TestDoWcount1(t *testing.T) {
    // test input to check for word frequency counting routines
    input_str := "Hello, this is me."
    
    // struct element order resembles sorted order
    output_tgt := []frame {
        {1, "Hello,"},
        {1, "is"},
        {1, "me."},
        {1, "this"},
    }
    
    // do the counting
    wcmap, wcmap_keys_sorted, _ := do_wcount(input_str,false)
    
    // get the counting output as a slice of structs
    output := do_print_freqs(wcmap,wcmap_keys_sorted,false)
    
    // check if the elements of the target and output structs match
    for i := 0; i <= len(output_tgt)-1; i++ {
        if output_tgt[i].wcount != output[i].wcount || output_tgt[i].word != output[i].word {
            t.Errorf(`Error matching count-word record (%d, '%s') = (%d, '%s')`, output_tgt[i].wcount, output_tgt[i].word, output[i].wcount, output[i].word)
        }
    }
}

func TestDoWcount2(t *testing.T) {
    // test input to check for word frequency counting routines
    input_str := "Hello, world!"
    
    // struct element order resembles sorted order
    output_tgt := []frame {
        {1, "Hello,"},
        {1, "world!"},
    }
    
    // do the counting
    wcmap, wcmap_keys_sorted, _ := do_wcount(input_str,false)
    
    // get the counting output as a slice of structs
    output := do_print_freqs(wcmap,wcmap_keys_sorted,false)
    
    for i := 0; i < len(output_tgt) ; i++ {
        if output_tgt[i].wcount != output[i].wcount || output_tgt[i].word != output[i].word {
            t.Errorf(`Error matching count-word record (%d, '%s') = (%d, '%s')`, output_tgt[i].wcount, output_tgt[i].word, output[i].wcount, output[i].word)
        }
    }
}
