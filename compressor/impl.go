package compressor

import (
	"bytes"
	"compress/gzip"
	"io"

	logger "github.com/sirupsen/logrus"
)

type Compressor struct {
	Buffer bytes.Buffer
}

func NewCompressor(rc io.ReadCloser) Store {
	// Prepare output writter
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

	return &Compressor{Buffer: buf}
}

func (c *Compressor) Read(p []byte) (n int, err error) {
	return c.Buffer.Read(p)
}
