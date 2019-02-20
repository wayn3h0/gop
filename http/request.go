package http

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/wayn3h0/gop/errors"
)

type requestBody struct {
	io.ReadCloser
	Parameters map[interface{}]interface{}
}

// Request represents a wrapper for http request.
// It stores the parameters directly in the http.Request by mutating the request.Body.
type Request struct {
	*http.Request
}

// SetParameter sets the parameter value.
func (r *Request) SetParameter(key interface{}, value interface{}) {
	body := r.Request.Body.(*requestBody)
	body.Parameters[key] = value
}

// Parameter get the parameter value.
func (r *Request) Parameter(key interface{}) interface{} {
	body := r.Request.Body.(*requestBody)
	return body.Parameters[key]
}

// Parameters returns the all parameters.
func (r *Request) Parameters() map[interface{}]interface{} {
	body := r.Request.Body.(*requestBody)
	return body.Parameters
}

// Method returns the http method.
// It resolves X-HTTP-Method-Override in header.
func (r Request) Method() string {
	method := r.Request.Header.Get("X-HTTP-Method-Override")
	if len(method) > 0 {
		return method
	}

	return r.Request.Method
}

// IPAddress returns the client ip address.
// It resolves X-Real-IP and X-Forwarded-For in header.
func (r Request) IPAddress() string {
	ip := r.Request.Header.Get("X-Real-IP")
	if len(ip) > 0 {
		return ip
	}

	ip = r.Request.Header.Get("X-Forwarded-For")
	if len(ip) > 0 {
		ip = strings.TrimSpace(strings.Split(ip, ",")[0])
		if len(ip) > 0 {
			return ip
		}
	}

	return r.Request.RemoteAddr
}

// ReadJSON reads body as JSON to obj.
func (r Request) ReadJSON(obj interface{}) error {
	buf, err := ioutil.ReadAll(r.Request.Body)
	if err != nil {
		return errors.Wrap(err, "http: could not read JSON data from request body")
	}

	if len(buf) == 0 {
		return nil
	}

	err = json.Unmarshal(buf, obj)
	if err != nil {
		return errors.Wrap(err, "http: could not unmarshal JSON data from request body")
	}

	return nil
}

// ReadXML reads body as XML to obj.
func (r Request) ReadXML(obj interface{}) error {
	buf, err := ioutil.ReadAll(r.Request.Body)
	if err != nil {
		return errors.Wrap(err, "http: could not read XML data from request body")
	}

	if len(buf) == 0 {
		return nil
	}

	err = xml.Unmarshal(buf, obj)
	if err != nil {
		return errors.Wrap(err, "http: could not unmarshal XML data from request body")
	}

	return nil
}

// ReadString reads body as string.
func (r Request) ReadString() (string, error) {
	buf, err := ioutil.ReadAll(r.Request.Body)
	if err != nil {
		return "", errors.Wrap(err, "http: could not read string data from request body")
	}

	return string(buf), nil
}

// NewRequest returns a new request.
func NewRequest(r *http.Request) *Request {
	if _, ok := r.Body.(*requestBody); !ok {
		body := &requestBody{
			ReadCloser: r.Body,
			Parameters: make(map[interface{}]interface{}),
		}
		r.Body = body
	}

	return &Request{
		Request: r,
	}
}
