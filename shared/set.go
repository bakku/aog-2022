package shared

type IntSet struct {
	Elements []int
}

func (set *IntSet) Contains(n int) bool {
	for _, element := range set.Elements {
		if element == n {
			return true
		}
	}

	return false
}

func (set *IntSet) Add(n int) {
	if !set.Contains(n) {
		set.Elements = append(set.Elements, n)
	}
}

func (set *IntSet) IsSubsetOf(other *IntSet) bool {
	for _, setElement := range set.Elements {
		currentElementContained := false

		for _, otherElement := range other.Elements {
			if setElement == otherElement {
				currentElementContained = true
			}
		}

		if !currentElementContained {
			return false
		}
	}

	return true
}

func (set *IntSet) Intersection(other *IntSet) *IntSet {
	result := []int{}

	for _, setElement := range set.Elements {
		for _, otherElement := range other.Elements {
			if setElement == otherElement {
				result = append(result, setElement)
			}
		}
	}

	return NewIntSet(result)
}

func (set *IntSet) Size() int {
	return len(set.Elements)
}

func NewIntSet(elements []int) *IntSet {
	s := &IntSet{
		Elements: make([]int, 0),
	}

	for _, n := range elements {
		s.Add(n)
	}

	return s
}
