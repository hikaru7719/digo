package writer

import (
	"fmt"
	"net"
	"strings"
)

type RequestWriter struct {
	domain   string
	header   []byte
	question []byte
	resolver string
}

func NewRequestWriter(domain string, resolver string) *RequestWriter {
	return &RequestWriter{domain: domain, resolver: resolver}
}

func (r *RequestWriter) CreateHeader() {
	header := make([]byte, 0, 12)

	header = append(header, 0x00)
	header = append(header, 0x00)

	header = append(header, 0x01)
	header = append(header, 0x00)

	header = append(header, 0x00)
	header = append(header, 0x01)

	for range make([]int, 6) {
		header = append(header, 0x00)
	}

	r.header = header
}

func (r *RequestWriter) CreateQuestion() {
	question := make([]byte, 0)
	split := strings.Split(r.domain, ".")
	for _, str := range split {
		question = append(question, byte(len(str)))
		for _, b := range []byte(str) {
			question = append(question, b)
		}
	}
	question = append(question, 0x00)

	question = append(question, 0x00)
	question = append(question, 0x01)
	question = append(question, 0x00)
	question = append(question, 0x01)
	r.question = question
}

func (r *RequestWriter) Dial() ([]byte, error) {
	conn, err := net.Dial("udp", r.resolver)
	defer conn.Close()
	if err != nil {
		fmt.Println("udp connection err")
		return nil, err
	}
	request := make([]byte, 0)
	request = append(request, r.header...)
	request = append(request, r.question...)

	_, err = conn.Write(request)
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, 1500)
	length, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	response := buffer[:length]

	return response, nil
}

func (r *RequestWriter) Run() ([]byte, error) {
	r.CreateHeader()
	r.CreateQuestion()
	return r.Dial()
}
