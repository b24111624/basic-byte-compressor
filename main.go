package main

import (
	"compress/gzip"
	"io"
	"os"
	"strings"

	logger "github.com/sirupsen/logrus"
)

func main() {
	// Init log
	logger.SetLevel((logger.DebugLevel))
	logger.Info("Starting basic-byte-compressor")

	// Prepare testing byte datas
	testData := "compressed data compressed data compressed data"

	rc := io.NopCloser(strings.NewReader(testData))
	defer rc.Close()

	// Prepare output file
	f, _ := os.Create("./gzip_file.gz")
	zw := gzip.NewWriter(f)
	defer zw.Close()

	// Compress data
	_, err := io.Copy(zw, rc)
	if err != nil {
		logger.WithField("err", err).Fatal("io.Copy failed")
	}
}
