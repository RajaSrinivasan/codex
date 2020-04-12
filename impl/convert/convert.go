package convert

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

var length vg.Length
var width vg.Length
var numMargin vg.Length
var txtStyle draw.TextStyle
var lineHeight vg.Length
var leftMargin vg.Length
var botMargin vg.Length

func init() {
	width, _ = vg.ParseLength("6.5in")
	numMargin, _ = vg.ParseLength("0.75in")
	txtStyle.Color = color.Black
	txtStyle.Font, _ = vg.MakeFont("Helvetica", 12.0)
	lineHeight = vg.Length(12.0)
	leftMargin = vg.Length(6.0)
	botMargin = vg.Length(6.0)
}

func paintLine(canvas draw.Canvas, idx int, line string, fromline int64, numlines int64) {
	linestripped := strings.Replace(line, "\t", "    ", -1)
	linenostr := fmt.Sprintf("%04d | %s", idx+int(fromline), linestripped)

	lineidx := numlines - int64(idx) - 1
	pt1 := vg.Point{X: 0.0 + leftMargin, Y: vg.Length(lineidx)*lineHeight + botMargin}
	canvas.FillText(txtStyle, pt1, linenostr)
	//fmt.Printf("%s\n", linenostr)
}

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

	length = lineHeight * vg.Length(linecount)
	pngcanvas, _ := draw.NewFormattedCanvas(width+2.0*leftMargin, length+2*botMargin, "png")
	canvas := draw.NewCanvas(pngcanvas, width, length)

	inp.Seek(0, 0)
	lineno := 0
	rdr := bufio.NewReader(inp)
	var line string

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

	paintLine(canvas, 0, line, linefrom, linecount)
	for lc := 1; int64(lc) < linecount; lc = lc + 1 {
		line, err = rdr.ReadString('\n')
		if err != nil {
			log.Printf("End of file before reading %d lines", linecount)
			return
		}
		paintLine(canvas, lc, line, linefrom, linecount)
	}
	ofilename := fmt.Sprintf("%s.%d.png", outbase, idx)
	ofile, _ := os.Create(ofilename)
	pngcanvas.WriteTo(ofile)
	ofile.Close()
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
