package csvreporter

import (
	"bytes"
	"encoding/csv"
	"fmt"

	"github.com/mclark4386/burgundy"
)

type CSVReporter struct {
	DontAddHeaders bool
}

func (r CSVReporter) Process(headers burgundy.Headers, rows []burgundy.Row) ([]byte, error) {
	data := [][]string{}
	buf := bytes.NewBuffer([]byte{})
	w := csv.NewWriter(buf)
	if !r.DontAddHeaders {
		data = append(data, headers)
	}
	for _, row := range rows {
		new_row := make([]string, len(row))
		for i, field := range row {
			if str, ok := field.(fmt.Stringer); ok {
				new_row[i] = str.String()
			} else {
				new_row[i] = fmt.Sprintf("%v", field)
			}
		}
		data = append(data, new_row)
	}
	if err := w.WriteAll(data); err != nil {
		return make([]byte, 0), err
	}
	return buf.Bytes(), nil
}
