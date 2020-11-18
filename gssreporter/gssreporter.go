package gssreporter

import (
	"io/ioutil"

	"github.com/mclark4386/burgundy"
	"gopkg.in/Iwark/spreadsheet.v2"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
)

type GSSReporter struct {
	CredFilePath  string
	CredFile      []byte
	TokenFilePath string
	SpreadSheetID string
	SheetID       uint
}

func (r GSSReporter) Process(headers burgundy.Headers, rows []burgundy.Row) ([]byte, error) {
	output_data := burgundy.DataBlockForProcessing(headers, rows, false)

	credFilePath := r.CredFilePath
	if credFilePath == "" {
		credFilePath = "credentials.json"
	}

	var data []byte
	var err error
	if r.CredFile == nil {
		data, err = ioutil.ReadFile(credFilePath)
		if err != nil {
			return []byte{}, err
		}
	} else {
		data = r.CredFile
	}

	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	if err != nil {
		return []byte{}, err
	}

	client := conf.Client(context.TODO())
	service := spreadsheet.NewServiceWithClient(client)

	spreadsheetID := r.SpreadSheetID
	ssheet, err := service.FetchSpreadsheet(spreadsheetID)
	if err != nil {
		return []byte{}, err
	}

	sheet, err := ssheet.SheetByID(r.SheetID)
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
