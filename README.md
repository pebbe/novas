A [Go](http://golang.org/) interface to the [Naval Observatory Vector Astrometry Software (NOVAS)](http://aa.usno.navy.mil/software/novas/).

Keywords: astronomy, astrometry, celestial mechanics, sun, moon, planets, stars

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

You can put the file in a different location. Then you need to provide
the full path at run time in the environment variable `JPLEPH`, or you
can set the full path at build time of the program that imports the
package, like this, with Go version 1.5 of newer:

    go build -ldflags "-X github.com/pebbe/novas.JPLephFile=/opt/opt_local/novas/JPLEPH" program.go

With Go version 1.4 or older:

    go build -ldflags "-X github.com/pebbe/novas.JPLephFile /opt/opt_local/novas/JPLEPH" program.go
