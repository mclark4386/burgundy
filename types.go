package burgundy

type Subject interface{}

type Headers []string

type Column struct {
	Name  string
	Index int
}

type Row []Item
type Item interface{}

type ColOrder []int

type FieldIndexing struct {
	Index    int
	TagIndex int
}

type FieldIndexes []FieldIndexing

// Len is part of sort.Interface.
func (f FieldIndexes) Len() int {
	return len(f)
}

// Swap is part of sort.Interface.
func (f FieldIndexes) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// Less is part of sort.Interface.
func (f FieldIndexes) Less(i, j int) bool {
	if f[i].TagIndex >= 0 && f[j].TagIndex >= 0 {
		return f[i].TagIndex < f[j].TagIndex
	} else if f[i].TagIndex >= 0 {
		return f[i].TagIndex <= f[j].Index
	} else if f[j].TagIndex >= 0 {
		return f[i].Index < f[j].TagIndex
	}
	return f[i].Index < f[j].Index
}
