package main

import (
	"compress/gzip"
	"io"
	"os"
	"strings"

	logger "github.com/sirupsen/logrus"

	"github.com/b24111624/basic-byte-compressor/compressor"
)

func main() {
	// Init log
	logger.SetLevel((logger.DebugLevel))
	logger.Info("Starting basic-byte-compressor")

	// Prepare testing byte datas
	testData := "compressed data compressed data compressed data"

	rc := io.NopCloser(strings.NewReader(testData))
	defer rc.Close()

	var b []byte
	compressorStore := compressor.NewCompressor(rc)
	//compressor.NewCompressor(rc)
	compressorStore.Read(b)

	// ungzip msg to check content
	zr, err := gzip.NewReader(compressorStore)
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
