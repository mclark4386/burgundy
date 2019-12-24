package burgundy

import "github.com/stretchr/testify/suite"

func RunTestReporter(e Reporter, expectedLen int, s suite.Suite) {
	if e != nil {
		data, err := e.Process([]string{}, []Row{})
		s.NoError(err)
		s.Equal(expectedLen, len(data))
	}
}
