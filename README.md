# Word frequency counting application in Go

This repository provides `wfreq`, a Go implementation of a word frequency counter.

It essentially simulates the Linux pipe

`
tr -s ' ' '\n' | sort | uniq -c | sort -n | tail
`

and is designed to generate exactly the same output.

## Building and running code

To build the code in this repo, follow these two steps:

``
$ go mod init && go build
``

In your folder holding a copy of this repo, you should see a `wfreq` binary file. Then, run this binary using:

``
$ ./wfreq
``

Enter some text and then press `return` and `Control+D` for the program to collect the input and making it unblock from expecting more input data.

Alternatively, if you need to avoid using `go build`, then just use

``
$ go run .
``

## Interacting with wfreq using IPC

`wfreq` can be used in two ways: (a) the first way is to directly invoke `wfreq` with the command `./wfreq`; and, (b) the second way is to pipe data from a command like `cat`.

To follow the second case, you can issue a command like

``
$ echo "Hello, world!" > myfile.txt
$ cat myfile.txt | ./wfreq
``

On my machine, the command above outputs

``
      1 Hello,
      1 is
      1 me
      1 this
``

Alternatively, you can combine `echo` with `./wfreq` via the command

``
$echo "Hello, brave world!" | ./wfreq
``

## Unit tests

I supply one file that provides a few toy tests for the routines involved in the implementation of `wfreq`. These tests are meant to provide some example invokations of the routines in `wfreq`, but are not essentially a least set of necessary tests that can verify that the original Linux pipe works correctly (see the top of this file).

To run unit tests, issue the command

``
$ go test -v
``

Note that the above command enables `verbose` output. Remove `-v` if you need more terse output.

## License

This repository is provided under the MIT license. It is authored by Sotiris Karavarsamis (s.karavarsamis@gmail.com).
