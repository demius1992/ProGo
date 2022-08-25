package main

import "io"

type CustomWriter struct {
	writer     io.Writer
	writeCount int
}

func NewCustomWriter(writer io.Writer) *CustomWriter {
	return &CustomWriter{writer, 0}
}

func (cw *CustomWriter) Write(slice []byte) (count int, err error) {
	count, err = cw.writer.Write(slice)
	cw.writeCount++
	Printfln("Custom Writer: %v bytes", count)
	return
}

func (cw *CustomWriter) Close() (err error) {
	if closer, ok := cw.writer.(io.Closer); ok {
		closer.Close()
	}
	Printfln("Total Writes: %v", cw.writeCount)
	return
}
