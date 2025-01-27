package pascals_triangle

var cache = map[int]([]int){}

func numbers(row int) []int {
	if row <= 1 {
		return []int{1}
	}
	if row == 2 {
		return []int{1, 1}
	}
	numbers := numbers(row - 1)
	slice, exists := cache[row]
	if exists {
		return slice
	}
	slice = []int{1}
	for index := 0; index < len(numbers)-1; index += 1 {
		pair := numbers[index : index+2]
		num := pair[0] + pair[1]
		slice = append(slice, num)
	}
	slice = append(slice, 1)
	cache[row] = slice
	return slice
}

/*
Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
*/
func generate(numRows int) [][]int {
	var pascalsTriangle = [][]int{}
	for row := 1; row <= numRows; row += 1 {
		pascalsTriangle = append(pascalsTriangle, numbers(row))
	}
	return pascalsTriangle
}
