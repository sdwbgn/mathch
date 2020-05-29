package mathch

import (
	"compress/gzip"
	"html/template"
	"io"
	"net/http"
	"strings"
)

func LoadHTML(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Cache-Control", "public, max-age=31536000")
	writer.Header().Set("ETag", ETag)
	if strings.HasPrefix(request.URL.Path, "/img/") {
		LoadImages(writer, request)
		return
	}
	switch request.URL.Path {
	case "/":
		tmpl, _ := template.ParseFiles("./static/index.html")
		if strings.Contains(request.UserAgent(), "Safari") && !strings.Contains(request.UserAgent(), "Chrome") {
			_ = tmpl.ExecuteTemplate(writer, "index", "jp2")
		} else {
			_ = tmpl.ExecuteTemplate(writer, "index", "webp")
		}
	case "/js/m.js":
		http.ServeFile(writer, request, "./js/m.js")
	case "/js/bootstrap.min.js":
		http.ServeFile(writer, request, "./js/bootstrap.min.js")
	case "/js/jquery-3.5.1.min.js":
		http.ServeFile(writer, request, "./js/jquery-3.5.1.min.js")
	case "/js/popper.min.js":
		http.ServeFile(writer, request, "./js/popper.min.js")
	case "/css/m.css":
		http.ServeFile(writer, request, "./css/m.css")
	case "/css/bootstrap.min.css":
		http.ServeFile(writer, request, "./css/bootstrap.min.css")
	case "/static/titillium-web-latin-ext.woff2":
		http.ServeFile(writer, request, "./static/titillium-web-latin-ext.woff2")
	case "/static/titillium-web-latin.woff2":
		http.ServeFile(writer, request, "./static/titillium-web-latin.woff2")
	default:
		http.NotFound(writer, request)
	}

}

func LoadImages(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Cache-Control", "public, max-age=31536000")
	writer.Header().Set("ETag", ETag)
	http.ServeFile(writer, request, "."+request.URL.Path)
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func Gzip(handle func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			handle(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzw := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		handle(gzw, r)
	}
}
