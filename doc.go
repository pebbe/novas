/*
Package novas provides an interface to the Naval Observatory Vector Astrometry Software (NOVAS) [astronomy, astrometry, celestial mechanics, sun, moon, planets, stars].

NOVAS homepage: http://aa.usno.navy.mil/software/novas/

You need a planetary ephemeris file JPLEPH. You can download it here: http://pkleiweg.home.xs4all.nl/jpleph/

Put it in this directory:

    $GOPATH/src/github.com/pebbe/novas/jpleph/

You can put the file in a different location. Then you need to provide
the full path at run time in the environment variable JPLEPH, or you
can set the full path at build time of the program that imports the
package, like this:

    go build -ldflags "-X github.com/pebbe/novas.JPLephFile=/opt/opt_local/novas/JPLEPH" program.go
*/
package novas
