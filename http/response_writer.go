package http

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/wayn3h0/gop/errors"
)

// ResponseWriter represents a wrapper for http.ResponseWriter.
type ResponseWriter struct {
	http.ResponseWriter
}

// WriteJSON writes the obj as JSON into the response body.
// It sets the Content-Type as "application/json".
func (r ResponseWriter) WriteJSON(obj interface{}) (int, error) {
	buf, err := json.Marshal(obj)
	if err != nil {
		return 0, errors.Wrap(err, "http: could not marshal entity to JSON data")
	}

	r.ResponseWriter.Header().Set("Content-Type", "application/json")

	n, err := r.ResponseWriter.Write(buf)
	if err != nil {
		return 0, errors.Wrap(err, "http: could not write JSON data to response")
	}

	return n, nil
}

// WriteXML writes the obj as XML into the response body.
// It sets the Content-Type as "application/xml".
func (r ResponseWriter) WriteXML(obj interface{}) (int, error) {
	buf, err := xml.Marshal(obj)
	if err != nil {
		return 0, errors.Wrap(err, "http: could not marshal entity to XML data")
	}

	// add xml header
	data := []byte(xml.Header)
	data = append(data, buf...)

	r.ResponseWriter.Header().Set("Content-Type", "application/xml")

	n, err := r.ResponseWriter.Write(data)
	if err != nil {
		return 0, errors.Wrap(err, "http: could not write XML data to response")
	}

	return n, nil
}

// WriteString writes the str into the response body.
// It sets the Content-Type as "text/plain".
func (r ResponseWriter) WriteString(str string) (int, error) {
	r.ResponseWriter.Header().Set("Content-Type", "text/plain")

	n, err := io.WriteString(r.ResponseWriter, str)
	if err != nil {
		return 0, errors.Wrap(err, "http: could not write string data to response")
	}

	return n, nil
}

// NewResponseWriter returns a new response writer.
func NewResponseWriter(rw http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: rw,
	}
}
