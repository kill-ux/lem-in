package lemIn

type DataStuct struct {
	Nants      int
	Start      string
	End        string
	Rooms      map[string][]int
	Links      [][]string
	Paths      [][]string
	PathGroups [][][]string
	Ants       [][]int
	Realtions  map[string][]string
	StepCalc   map[int]string
}

var Data = &DataStuct{
	Rooms:     make(map[string][]int),
	Realtions: make(map[string][]string),
	StepCalc:  make(map[int]string),
}
