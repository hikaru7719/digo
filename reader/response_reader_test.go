package reader

import "testing"

func TestResponseReader_ReadHeader(t *testing.T) {
	cases := map[string]struct {
		reader          *ResponseReader
		expectARecord   uint16
		expectNSRecord  uint16
		expectAddRecord uint16
	}{
		"success": {
			reader:          testNewResponseReader(t),
			expectARecord:   1,
			expectNSRecord:  0,
			expectAddRecord: 0,
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			tc.reader.ReadHeader()
			if tc.reader.aRecord != tc.expectARecord {
				t.Errorf("want %d, but actual %d\n", tc.expectARecord, tc.reader.aRecord)
			}

			if tc.reader.nsRecord != tc.expectNSRecord {
				t.Errorf("want %d, but actual %d\n", tc.expectNSRecord, tc.reader.nsRecord)
			}

			if tc.reader.addRecord != tc.expectAddRecord {
				t.Errorf("want %d, but actual %d\n", tc.expectAddRecord, tc.reader.addRecord)
			}
		})
	}
}

func TestResponseReader_ReadARecord(t *testing.T) {
	cases := map[string]struct {
		reader *ResponseReader
	}{
		"success": {
			reader: testNewResponseReader(t),
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			tc.reader.ReadARecord()
		})
	}
}

func testNewResponseReader(t *testing.T) *ResponseReader {
	testResponse := []byte{
		0x00, 0x00, 0x81, 0x80, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x03, 0x77, 0x77, 0x77,
		0x04, 0x6a, 0x70, 0x72, 0x73, 0x02, 0x63, 0x6f, 0x02, 0x6a, 0x70, 0x00, 0x00, 0x01, 0x00, 0x01,
		0xc0, 0x0c, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x01, 0x2c, 0x00, 0x04, 0x75, 0x68, 0x85, 0xa5,
	}
	testQuery := []byte{
		0x03, 0x77, 0x77, 0x77,
		0x04, 0x6a, 0x70, 0x72, 0x73, 0x02, 0x63, 0x6f, 0x02, 0x6a, 0x70, 0x00, 0x00, 0x01, 0x00, 0x01,
	}
	return NewResponseReader(testQuery, testResponse)
}
