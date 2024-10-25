package sorting

type Sorter interface {
	Sort([]int) []int
	Name() string
}
