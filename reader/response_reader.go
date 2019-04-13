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
	aEnd        int
	nsEnd       int
	addEnd      int
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

func (r *ResponseReader) ReadARecord() {
	var current int
	current += headSize + len(r.query)
	offset := r.response[current : current+2]
	fmt.Println(offset)
	current += 2

	resourceType := r.response[current : current+2]
	fmt.Printf("resourceType: %X\n", resourceType)
	current += 2

	classType := r.response[current : current+2]
	fmt.Printf("class: %X\n", classType)
	current += 2

	ttl := r.response[current : current+4]
	fmt.Printf("ttl: %d\n", binary.BigEndian.Uint16(ttl))
	current += 4

	length := r.response[current : current+2]
	fmt.Printf("length: %d byte\n", binary.BigEndian.Uint16(length))
	intLength := int(binary.BigEndian.Uint16(length))
	current += 2

	answer := r.response[current : current+intLength]
	if intLength == 4 {
		a := int(answer[0])
		b := int(answer[1])
		c := int(answer[2])
		d := int(answer[3])
		fmt.Printf("ipAddress: %d.%d.%d.%d\n", a, b, c, d)
	}
}
