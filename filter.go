package webassets

import (
	"bytes"
	"io"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"

	libsass "github.com/wellington/go-libsass"
)

func (w *webassets) libscss(r io.Reader, paths ...string) (io.Reader, error) {
	var b bytes.Buffer

	comp, err := libsass.New(&b, r, libsass.IncludePaths(paths))
	if err != nil {
		return nil, err
	}
	if err := comp.Run(); err != nil {
		return nil, err
	}
	return &b, nil
}

func (w *webassets) cssmin(r io.Reader) (io.Reader, error) {
	var b bytes.Buffer

	m := minify.New()
	m.AddFunc("css", css.Minify)

	if err := m.Minify("css", &b, r); err != nil {
		return nil, err
	}
	return &b, nil
}

func (w *webassets) jsmin(r io.Reader) (io.Reader, error) {
	var b bytes.Buffer

	m := minify.New()
	m.AddFunc("js", js.Minify)

	if err := m.Minify("js", &b, r); err != nil {
		return nil, err
	}
	return &b, nil
}

func (w *webassets) htmlmin(r io.Reader) (io.Reader, error) {
	var b bytes.Buffer

	m := minify.New()
	m.AddFunc("html", html.Minify)

	if err := m.Minify("html", &b, r); err != nil {
		return nil, err
	}
	return &b, nil
}
