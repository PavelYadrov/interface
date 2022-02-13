package types

type Something struct {
	Slice []int
}

func (smth Something) Len() int {
	return len(smth.Slice)
}

func (is Something) Swap(i, j int) {
	is.Slice[i], is.Slice[j] = is.Slice[j], is.Slice[i]
}
