package flatten_nested_list_iterator

type NestedIntegerInterface interface {
	IsInteger() bool
	GetInteger() int
	SetInteger(value int)
	Add(elem NestedInteger)
	GetList() []*NestedInteger
}

type NestedInteger = struct {
	NestedIntegerInterface
}

type NestedIterator struct {
	stack   [][]*NestedInteger // Stack to store lists that need to be processed
	current *NestedInteger     // Current element being processed
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		stack: [][]*NestedInteger{nestedList},
	}
}

func (this *NestedIterator) Next() int {
	// Since HasNext() ensures we have a valid integer,
	// we can safely return the current integer
	if this.HasNext() {
		val := this.current.GetInteger()
		this.current = nil
		return val
	}
	return 0 // Should never reach here if used correctly
}

func (this *NestedIterator) HasNext() bool {
	// If we already have a current integer, return true
	if this.current != nil {
		return true
	}

	// Process until we find the next integer
	for len(this.stack) > 0 {
		// Get the current list from top of stack
		currentList := this.stack[len(this.stack)-1]

		// If current list is empty, pop it and continue
		if len(currentList) == 0 {
			this.stack = this.stack[:len(this.stack)-1]
			continue
		}

		// Get the first element from current list
		element := currentList[0]
		// Update the current list by removing the first element
		this.stack[len(this.stack)-1] = currentList[1:]

		if element.IsInteger() {
			// If it's an integer, set it as current and return true
			this.current = element
			return true
		} else {
			// If it's a list, push it to stack
			list := element.GetList()
			if len(list) > 0 {
				this.stack = append(this.stack, list)
			}
		}
	}

	return false
}
