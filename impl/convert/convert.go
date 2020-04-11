package convert

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convertFragment(inp *os.File, outbase string, idx int, frag string) {
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

	inp.Seek(0, 0)
	lineno := 0
	rdr := bufio.NewReader(inp)
	var line string
	//var err error

	for {
		line, err = rdr.ReadString('\n')
		if err != nil {
			log.Printf("End of File before line %d", linefrom)
			return
		}
		lineno = lineno + 1
		if int64(lineno) >= linefrom {
			break
		}
	}
	for lc := 0; int64(lc) < linecount; lc = lc + 1 {
		fmt.Printf("%04d | %s", linefrom+int64(lc), line)
		line, err = rdr.ReadString('\n')
		if err != nil {
			log.Printf("End of file before reading %d lines", linecount)
			return
		}

	}
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
