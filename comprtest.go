package main

import (
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

type CompressMethod int

const (
	CompressGzip = iota
	CompressLz4
)

type LogLevelType int

const (
	LogError = iota
	LogWarn
	LogInfo
)

var pageSize = 16384

var terse bool
var maxPages int64
var logLevel LogLevelType
var inFileName string
var comprMethod CompressMethod
var printHeader bool

//
// If there is an error, print the msg string and then panic.
//
func checkError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s\n", msg)
		panic(err)
	}
}

//
// Given a page of data, compress it, and return the compressed length and
// an error indication. The compressed data is thrown away.
//
func comprPage(data []byte) (int, error) {
	var wBuf bytes.Buffer
	var zw *gzip.Writer
	var retLen int

	zw = gzip.NewWriter(&wBuf)

	_, err := zw.Write(data)
	if err != nil {
		return 0, err
	}

	if err := zw.Close(); err != nil {
		return 0, err
	}

	retLen = wBuf.Len()

	return retLen, err
}

//
// Do various initializations
//
func doInit() {
	rand.Seed(time.Now().Unix())
}

//
// Convert given number to next higher power of 2 and return it.
//
func toPower2(num uint) uint {
	num--

	num |= num >> 1
	num |= num >> 2
	num |= num >> 4
	num |= num >> 8
	num |= num >> 16

	num++

	return num
}

//
// Echo the input args for debugging
//
func echoArgs() {
	fmt.Printf("Page Size: %d\n", pageSize)
	fmt.Printf("Compress method: %d\n", comprMethod)
}

//
// Parse the command line arguments
//
// Usage: comprtst -pgsz <pagesize> <inputfile>
//
func parseArgs() {
	var pgSzp = flag.Int("pgsz", 8192, "page size")
	var cMethdp = flag.Int("cmethod", 0, "compression method")
	var terFlagp = flag.Bool("terse", true, "terse output")
	var lgLvlp = flag.Int("loglevel", 0, "log level")
	var pHdrp = flag.Bool("h", false, "print header")

	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		panic("Bad args")
	}

    inFileName = flag.Arg(0)

	if *cMethdp != int(CompressGzip) && *cMethdp != int(CompressLz4) {
		panic("Wrong compression method supplied")
	}

	pageSize = *pgSzp

	comprMethod = CompressMethod(*cMethdp)

	terse = *terFlagp

	logLevel = LogLevelType(*lgLvlp)

	printHeader = *pHdrp

	if !terse {
		echoArgs()
	}
}

func main() {
	var fName1 string
	var pageCount int64
	var totalComprSz int64

	startTime := time.Now()

	parseArgs()

	if !terse {
		fmt.Printf("Start time: %v\n\n",
			startTime.Format("2016-01-02 15:04:05"))
	}

	doInit()

	fName1 = flag.Arg(0)

	fId1, err := os.Open(fName1)
	checkError(err, "Open error")

	fInfo, err := fId1.Stat()
	checkError(err, "File Stat")

	fSz := fInfo.Size()

	data := make([]byte, pageSize)

	for {
		count, err := fId1.Read(data)
		if count == 0 || err == io.EOF {
			break
		}
		checkError(err, "Read error")

		comprCount, err := comprPage(data)
		checkError(err, "comprPage error")

		totalComprSz += int64(comprCount)

		pageCount++
	}

	endTime := time.Now()

	cRatio := float64(fSz) / float64(totalComprSz)

	if printHeader {
		fmt.Printf("%12s %12s %12s %12s %12s\n", "PageSz", "Pages", "Compressed", "Uncompressed", "Ratio")
	}

	fmt.Printf("%12d %12d %12d %12d %10.2f\n", pageSize, pageCount, totalComprSz, fSz, cRatio)

	if !terse {
		fmt.Printf("\nEnd time: %v\n", endTime.Format("2016-01-02 15:04:05"))
		fmt.Printf("Elapsed time: %v\n", endTime.Sub(startTime))
	}
}
