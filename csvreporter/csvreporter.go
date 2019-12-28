package csvreporter

import (
	"bytes"
	"encoding/csv"

	"github.com/mclark4386/burgundy"
)

type CSVReporter struct {
	DontAddHeaders bool
}

func (r CSVReporter) Process(headers burgundy.Headers, rows []burgundy.Row) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	w := csv.NewWriter(buf)
	data := burgundy.DataBlockForProcessing(headers, rows, r.DontAddHeaders)
	if err := w.WriteAll(data); err != nil {
		return make([]byte, 0), err
	}
	return buf.Bytes(), nil
}
