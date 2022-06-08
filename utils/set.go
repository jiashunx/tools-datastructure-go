package utils

/**
 * 使用空struct{}实现set
 */
type StringSet map[string]struct{}

func NewStringSet() StringSet {
	return make(StringSet)
}

func (o StringSet) Has(arg string) bool {
	_, ok := o[arg]
	return ok
}

func (o StringSet) Add(arg string) {
	o[arg] = struct{}{}
}

func (o StringSet) Remove(arg string) {
	delete(o, arg)
}

func (o StringSet) Size() int {
	return len(o)
}

func (o StringSet) IsEmpty() bool {
	return o.Size() == 0
}
