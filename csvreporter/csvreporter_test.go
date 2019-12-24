package csvreporter

import (
	"testing"

	"github.com/mclark4386/burgundy"
	"github.com/stretchr/testify/suite"
)

type Lamp struct {
	Bulbs    int    `col:"1"`
	Style    string `col:"0"`
	Watts    float64
	Password string `col:"-"`
	Owner    Anchor
}

type Anchor struct {
	Name    string
	Station string
}

func (a Anchor) String() string {
	if a.Name == "" && a.Station == "" {
		return "wooo I'm a spooky ghost!"
	}
	return "This is " + a.Name + ", " + a.Station + " news."
}

type CSVReporterTestSuite struct {
	suite.Suite
}

func TestCSVReporterTestSuite(t *testing.T) {
	suite.Run(t, new(CSVReporterTestSuite))
}

func (cs *CSVReporterTestSuite) TestCSVReporter() {
	i_love := []Lamp{
		{
			Bulbs:    2,
			Style:    "70s",
			Watts:    60,
			Password: "SexPanther",
			Owner: Anchor{
				Name:    "Ron Burgundy",
				Station: "Channel 4",
			},
		},
		{
			Bulbs:    3,
			Style:    "70s",
			Watts:    45,
			Password: "ThatEscalatedQuickly",
		},
	}
	expectedLen := 107 //honestly had to run it and visually verify to get this number... you have been warned

	data, err := burgundy.Process(i_love, CSVReporter{})

	print(string(data))
	cs.Equal(expectedLen, len(data))
	cs.NoError(err)
}
