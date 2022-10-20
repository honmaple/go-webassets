package webassets

import (
	"bytes"
	"io"
	"io/ioutil"
	"sync"
)

type webassets struct {
	buffer sync.Pool
}

func (w *webassets) Run(files []string, filters []string) (io.Reader, error) {
	var b bytes.Buffer
	for _, file := range files {
		buf, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		b.Write(buf)
		b.WriteString("\n")
	}
	var (
		err error
		r   io.Reader
	)
	r = &b
	for _, filter := range filters {
		switch filter {
		case "libscss":
			r, err = w.libscss(r)
		case "cssmin":
			r, err = w.cssmin(r)
		case "jsmin":
			r, err = w.jsmin(r)
		case "htmlmin":
			r, err = w.htmlmin(r)
		}
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

func New() *webassets {
	return &webassets{buffer: sync.Pool{
		New: func() interface{} {
			var b bytes.Buffer
			return &b
		},
	}}
}
