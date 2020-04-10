package convert

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func convertFragment(rdr io.Reader, outbase string, idx int, frag string) {
	linespec := strings.Split(frag, ":")
	if len(linespec) != 2 {
		log.Printf("Invalid fragment spec %s", frag)
		return
	}
	linefrom, err := strconv.ParseInt(linespec[0], 10, 16)
	if err != nil {
		log.Printf("Invalid line number %s", linespec[0])
		return
	}
	linecount, err := strconv.ParseInt(linespec[1], 10, 16)
	if err != nil {
		log.Printf("Invalid line count %s", linespec[1])
	}

	log.Printf("Will create %s.%d.png with lines %d to %d", outbase, idx, linefrom, linefrom+linecount-1)
}

func Convert(inp string, outbase string, frags []string) {
	log.Printf("Extracting fragments from %s into graphic files %s", inp, outbase)
	inpfile, err := os.Open(inp)
	if err != nil {
		log.Fatal(err)
	}
	defer inpfile.Close()

	for idx, frag := range frags {
		convertFragment(inpfile, outbase, idx, frag)
	}
}
