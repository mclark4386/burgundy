package burgundy

type Reporter interface {
	Process(Headers, []Row) ([]byte, error)
}
