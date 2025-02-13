package course_schedule

/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	prerequisitesMap := map[int]([]int){}
	statusMap := map[int]int{}
	const (
		unvisited = 0
		visiting  = 1
		visited   = 2
	)
	for _, pair := range prerequisites {
		course, prerequisite := pair[0], pair[1]
		if _, exists := prerequisitesMap[course]; !exists {
			prerequisitesMap[course] = []int{}
			statusMap[course] = unvisited
		}
		prerequisitesMap[course] = append(prerequisitesMap[course], prerequisite)
	}
	var dfs func(course int) bool
	dfs = func(course int) bool {
		if statusMap[course] == visiting {
			return true
		}
		if statusMap[course] == visited {
			return false
		}
		statusMap[course] = visiting
		for _, prerequisite := range prerequisitesMap[course] {
			if dfs(prerequisite) {
				return true
			}
		}
		statusMap[course] = visited
		return false
	}
	for course := range numCourses {
		conflict := dfs(course)
		if conflict {
			return false
		}
	}
	return true
}
