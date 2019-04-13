package writer

import (
	"reflect"
	"testing"
)

func TestRequestWriter_CreateHeader(t *testing.T) {
	cases := map[string]struct {
		writer       *RequestWriter
		expectHeader []byte
	}{
		"success": {
			writer:       testNewRequestWriter(),
			expectHeader: createExpectHeader(),
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			tc.writer.CreateHeader()
			if !reflect.DeepEqual(tc.writer.header, tc.expectHeader) {
				t.Errorf("want %v, but actual %v\n", tc.expectHeader, tc.writer.header)
			}
		})
	}
}

func TestRequestWriter_CreateQuestion(t *testing.T) {
	cases := map[string]struct {
		writer         *RequestWriter
		expectQuestion []byte
	}{
		"success": {
			writer:         testNewRequestWriter(),
			expectQuestion: createExpectQuestion(),
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			tc.writer.CreateQuestion()
			if !reflect.DeepEqual(tc.writer.question, tc.expectQuestion) {
				t.Errorf("want %v, but actual %v\n", tc.expectQuestion, tc.writer.question)
			}
		})
	}
}

func TestRequestWriter_Run(t *testing.T) {
	cases := map[string]struct {
		writer         *RequestWriter
		expectResponse []byte
	}{
		"success": {
			writer:         testNewRequestWriter(),
			expectResponse: createResponse(),
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			response, err := tc.writer.Run()
			if err != nil {
				t.Error(err)
			}

			if reflect.DeepEqual(response, tc.expectResponse) {
				t.Errorf("want %v\n, but actual %v\n", tc.expectResponse, response)
			}
		})
	}
}

func testNewRequestWriter() *RequestWriter {
	return NewRequestWriter("www.jprs.co.jp", "103.5.140.1:53")
}

func createExpectHeader() []byte {
	header := []byte{0x00, 0x00, 0x01, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	return header
}

func createExpectQuestion() []byte {
	question := []byte{0x03, 0x77, 0x77, 0x77, 0x04, 0x6a, 0x70, 0x72, 0x73, 0x02, 0x63, 0x6f, 0x02, 0x6a, 0x70, 0x00, 0x00, 0x01, 0x00, 0x01}
	return question
}

func createResponse() []byte {
	return []byte{
		0x00, 0x00, 0x81, 0x80, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x03, 0x77, 0x77, 0x77,
		0x04, 0x6a, 0x70, 0x72, 0x73, 0x02, 0x63, 0x6f, 0x02, 0x6a, 0x70, 0x00, 0x00, 0x01, 0x00, 0x01,
		0xc0, 0x0c, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x01, 0x2c, 0x00, 0x04, 0x75, 0x68, 0x85, 0xa5,
	}
}
