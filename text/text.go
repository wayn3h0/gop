package text

import (
	"io"
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func transformEncoding(rawReader io.Reader, trans transform.Transformer) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(rawReader, trans))
	if err == nil {
		return string(ret), nil
	} else {
		return "", err
	}
}

// Convert a string encoding from ShiftJIS to UTF-8
func FromShiftJIS(str string) (string, error) {
	return transformEncoding(strings.NewReader(str), japanese.ShiftJIS.NewDecoder())
}

// Convert a string encoding from UTF-8 to ShiftJIS
func ToShiftJIS(str string) (string, error) {
	return transformEncoding(strings.NewReader(str), japanese.ShiftJIS.NewEncoder())
}

// Convert a string encoding from GBK to UTF-8
func FromGBK(str string) (string, error) {
	return transformEncoding(strings.NewReader(str), simplifiedchinese.GBK.NewDecoder())
}

// Convert a string encoding from UTF-8 to GBK
func ToGBK(str string) (string, error) {
	return transformEncoding(strings.NewReader(str), simplifiedchinese.GBK.NewEncoder())
}
