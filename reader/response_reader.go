package reader

import (
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
)

// responseのheaderサイズは12byte
const headSize = 12

type ResponseReader struct {
	query       []byte
	queryLength int
	response    []byte
	header      []byte
	aRecord     uint16
	nsRecord    uint16
	addRecord   uint16
}

func NewResponseReader(query, response []byte) *ResponseReader {
	return &ResponseReader{query: query, queryLength: len(query), response: response}
}

func (r *ResponseReader) GetHeader() {
	r.header = r.response[:headSize]
}
func (r *ResponseReader) ReadHeader() error {
	r.GetHeader()
	var current int
	id := r.header[:2]
	fmt.Println(id)
	current += 2
	errFlag := r.header[current : current+2]
	if !reflect.DeepEqual(errFlag, []byte{0x81, 0x80}) {
		fmt.Println("error flag")
		return errors.New("invalid flag")
	}
	current += 2
	qCount := r.header[current : current+2]
	fmt.Println(binary.BigEndian.Uint16(qCount))
	current += 2
	aCount := r.header[current : current+2]
	fmt.Println(binary.BigEndian.Uint16(aCount))
	r.aRecord = binary.BigEndian.Uint16(aCount)
	current += 2

	nsCount := r.header[current : current+2]
	fmt.Println(binary.BigEndian.Uint16(nsCount))
	r.nsRecord = binary.BigEndian.Uint16(nsCount)
	current += 2

	addCount := r.header[current : current+2]
	fmt.Println(binary.BigEndian.Uint16(addCount))
	r.addRecord = binary.BigEndian.Uint16(addCount)
	current += 2
	return nil
}
