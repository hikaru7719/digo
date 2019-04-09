package reader

import (
	"errors"
	"fmt"
	"reflect"
)

// responseのheaderサイズは12byte
const headSize = 12

type ResponseReader struct {
	query     []byte
	response  []byte
	header    []byte
	aRecord   int
	nsRecord  int
	addRecord int
}

func NewResponseReader(query, packet []byte) *ResponseReader {
	return &ResponseReader{query: query, response: packet}
}

func (r *ResponseReader) GetHeader() {
	r.header = r.response[:headSize]
}
func (r *ResponseReader) ReadHeader() error {
	var current int
	_ := r.header[:2]
	current += 2
	errFlag := r.header[current : current+2]
	if reflect.DeepEqual(errFlag, []byte{0x81, 0x80}) {
		fmt.Println("error flag")
		return errors.New("invalid flag")
	}
	current += 2
	aCount := r.header[current : current+2]

	return nil
}
