package main

import (
	"bytes"
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
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	// Compress data to buffer
	_, err := io.Copy(zw, rc)
	if err != nil {
		logger.WithField("err", err).Fatal("io.Copy failed")
	}

	if err := zw.Close(); err != nil {
		logger.WithField("err", err).Fatal("zw.close failed")
	}

	zr, err := gzip.NewReader(&buf)
	if err != nil {
		logger.WithField("err", err).Fatal("gzip.NewReader failed")
	}

	if _, err := io.Copy(os.Stdout, zr); err != nil {
		logger.WithField("err", err).Fatal("io.Copy to stdout failed")
	}

	if err := zr.Close(); err != nil {
		logger.WithField("err", err).Fatal("zr.Close failed")
	}

}
