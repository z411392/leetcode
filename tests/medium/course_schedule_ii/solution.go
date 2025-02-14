package course_schedule_ii

import "fmt"

/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.
*/
func findOrder(numCourses int, prerequisites [][]int) []int {
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
	courses := []int{}
	var dfs func(course int) error
	dfs = func(course int) error {
		if statusMap[course] == visiting {
			return fmt.Errorf("visiting")
		}
		if statusMap[course] == visited {
			return nil
		}
		statusMap[course] = visiting
		for _, prerequisite := range prerequisitesMap[course] {
			err := dfs(prerequisite)
			if err != nil {
				return err
			}
		}
		courses = append(courses, course)
		statusMap[course] = visited
		return nil
	}
	for course := range numCourses {
		err := dfs(course)
		if err != nil {
			return []int{}
		}
	}
	return courses

}
