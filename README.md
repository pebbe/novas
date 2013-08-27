A [Go](http://golang.org/) interface to the [Naval Observatory Vector Astrometry Software (NOVAS)](http://aa.usno.navy.mil/software/novas/).

## Docs

 * [package help](http://godoc.org/github.com/pebbe/novas)
 * [examples](https://github.com/pebbe/novas/tree/master/examples)

## Installation

To install the package itself, run:

    go get github.com/pebbe/novas

To run programs using this package, you will also need a planetary
ephemeris file, called `JPLEPH`. To create this file, download the files
from this site:

ftp://ssd.jpl.nasa.gov/pub/eph/planets/fortran/

and follow the instructions in `userguide.txt`
