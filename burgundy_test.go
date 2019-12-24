package burgundy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type BurgundyTestSuite struct {
	suite.Suite
}

type Lamp struct {
	Bulbs    int    `col:"1"`
	Style    string `col:"0"`
	Watts    float64
	Password string `col:"-"`
}

var expectedLampHeaders = []string{
	"Style",
	"Bulbs",
	"Watts",
}

var expectedLampOrder = []int{1, 0, 2}

type Emotions struct {
	Complex   int `col:"1"`
	Panther   float32
	Style     string `col:"0"`
	Watts     float64
	Password  string   `col:"-"`
	GlassCage []string `col:"3"`
}

var expectedEmotionOrder = []int{2, 0, 1, 5, 3}

func TestBurgundyTestSuite(t *testing.T) {
	suite.Run(t, new(BurgundyTestSuite))
}

func (bs *BurgundyTestSuite) TestProcessRow() {
	i_love := Lamp{
		Bulbs: 2,
		Style: "70s",
		Watts: 60,
	}

	expectedLampValues := []interface{}{
		"70s",
		2,
		float64(60),
	}

	test_row := processRow(i_love, expectedLampOrder)

	fmt.Printf("test_row: %+v\n", test_row)
	bs.Equal(3, len(test_row))
	for i, expectedValue := range expectedLampValues {
		bs.Equal(expectedValue, test_row[i])
	}

	test_row = processRow(&i_love, expectedLampOrder)

	fmt.Printf("test_row: %+v\n", test_row)
	bs.Equal(3, len(test_row))
	for i, expectedValue := range expectedLampValues {
		bs.Equal(expectedValue, test_row[i])
	}
}

func (bs *BurgundyTestSuite) TestProcessHeaders() {
	i_love := Lamp{
		Bulbs: 2,
		Style: "70s",
		Watts: 60,
	}

	test_headers := processHeaders(i_love, expectedLampOrder)

	fmt.Printf("test_headers: %+v\n", test_headers)
	bs.Equal(3, len(test_headers))
	for i, header := range expectedLampHeaders {
		bs.Equal(header, test_headers[i])
	}
}

func (bs *BurgundyTestSuite) TestProcessOrder() {
	i_love := Lamp{
		Bulbs: 2,
		Style: "70s",
		Watts: 60,
	}

	test_order := processOrder(i_love)

	fmt.Printf("test_order: %+v\n", test_order)

	bs.Equal(len(expectedLampOrder), len(test_order))
	for i, expectedIndex := range expectedLampOrder {
		bs.Equal(expectedIndex, test_order[i])
	}

	test_order = processOrder(Emotions{})

	fmt.Printf("test_order2: %+v\n", test_order)

	bs.Equal(len(expectedEmotionOrder), len(test_order))
	for i, expectedIndex := range expectedEmotionOrder {
		bs.Equal(expectedIndex, test_order[i])
	}
}

func (bs *BurgundyTestSuite) TestProcess() {
	i_love := []Lamp{
		{
			Bulbs: 2,
			Style: "70s",
			Watts: 60,
		},
		{
			Bulbs: 3,
			Style: "70s",
			Watts: 45,
		},
	}
	expectedLen := 0

	data, err := Process(i_love, MockReporter{TestSuite: bs.Suite})

	bs.Equal(expectedLen, len(data))
	bs.NoError(err)

	data, err = Process([]Lamp{}, MockReporter{TestSuite: bs.Suite})

	bs.Equal(expectedLen, len(data))
	bs.NoError(err)
}

func (bs *BurgundyTestSuite) TestProcess_shouldFailWithWrongType() {
	i_love := Lamp{
		Bulbs: 2,
		Style: "70s",
		Watts: 60,
	}
	expectedLen := 0

	data, err := Process(i_love, MockReporter{TestSuite: bs.Suite})

	bs.Equal(expectedLen, len(data))
	bs.Error(err)

	data, err = Process(4, MockReporter{TestSuite: bs.Suite})

	bs.Equal(expectedLen, len(data))
	bs.Error(err)
}
