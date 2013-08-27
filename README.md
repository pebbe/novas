A [Go](http://golang.org/) interface to the [Naval Observatory Vector Astrometry Software (NOVAS)](http://aa.usno.navy.mil/software/novas/).

## Docs

 * [package help](http://godoc.org/github.com/pebbe/novas)
 * [examples](https://github.com/pebbe/novas/tree/master/examples)

## Installation

To install the package itself, run:

    go get github.com/pebbe/novas

You also need a planetary ephemeris file `JPLEPH`. You can download it from here:

http://pkleiweg.home.xs4all.nl/jpleph/

Put the file in this directory:

    $GOPATH/src/github.com/pebbe/novas/jpleph/
