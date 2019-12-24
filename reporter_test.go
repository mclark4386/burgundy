package burgundy

import "github.com/stretchr/testify/suite"

type MockReporter struct {
	TestSuite suite.Suite
}

func (r MockReporter) Process(headers Headers, objs []Row) ([]byte, error) {
	for _, row := range objs {
		r.TestSuite.Equal(len(headers), len(row))
	}
	return make([]byte, 0), nil
}

func (bs *BurgundyTestSuite) TestMockReporter() {
	expectedLen := 0
	me := MockReporter{TestSuite: bs.Suite}
	RunTestReporter(me, expectedLen, bs.Suite)
}
