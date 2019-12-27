package gssreporter

import (
	"fmt"
	"io/ioutil"

	"github.com/mclark4386/burgundy"
	"gopkg.in/Iwark/spreadsheet.v2"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
)

type GSSReporter struct {
	CredFilePath  string
	TokenFilePath string
	SpreadSheetID string
	SheetID       uint
}

func (r GSSReporter) Process(headers burgundy.Headers, rows []burgundy.Row) ([]byte, error) {
	output_data := [][]string{}
	output_data = append(output_data, headers)
	for _, row := range rows {
		new_row := make([]string, len(row))
		for i, field := range row {
			if str, ok := field.(fmt.Stringer); ok {
				new_row[i] = str.String()
			} else {
				new_row[i] = fmt.Sprintf("%v", field)
			}
		}
		output_data = append(output_data, new_row)
	}

	credFilePath := r.CredFilePath
	if credFilePath == "" {
		credFilePath = "credentials.json"
	}

	data, err := ioutil.ReadFile(credFilePath)
	if err != nil {
		return []byte{}, err
	}

	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	if err != nil {
		return []byte{}, err
	}

	client := conf.Client(context.TODO())
	service := spreadsheet.NewServiceWithClient(client)

	spreadsheetID := r.SpreadSheetID
	spreadsheet, err := service.FetchSpreadsheet(spreadsheetID)
	if err != nil {
		return []byte{}, err
	}

	sheet, err := spreadsheet.SheetByID(r.SheetID)
	if err != nil {
		return []byte{}, err
	}

	for r, row := range output_data {
		for c, field := range row {
			sheet.Update(r, c, field)
		}
	}

	err = sheet.Synchronize()
	if err != nil {
		return []byte{}, err
	}

	return []byte{}, nil
}
