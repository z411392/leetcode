package gas_station

/*
There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].

You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith station to its next (i + 1)th station. You begin the journey with an empty tank at one of the gas stations.

Given two integer arrays gas and cost, return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1. If there exists a solution, it is guaranteed to be unique.
*/
func canCompleteCircuit(gas []int, cost []int) int {
	positions := len(gas)
	if positions == 0 {
		return -1
	}
	if positions == 1 {
		// 要能夠出發
		if gas[0] <= 0 {
			return -1
		}
		// 要能夠回到自己
		if (gas[0] - cost[0]) < 0 {
			return -1
		}
		return 0
	}
	diffs := []int{}
	for i := 0; i < positions; i += 1 {
		diff := gas[i] - cost[i]
		diffs = append(diffs, diff)
	}
	for i := 0; i < positions; i += 1 {
		if diffs[i] <= 0 {
			continue
		}
		sum := 0
		traversed := 0
		found := false
		for {
			traversed += 1
			next := (i + traversed - 1) % positions
			sum += diffs[next]
			// fmt.Printf("i=%v, next=%v, traversed=%v, positions=%v, sum=%v, \n", i, next, traversed, positions, sum)
			if traversed == positions {
				// 已走完全部節點時能源不小於 0 即可（小於 0 表示無法到達）
				if sum < 0 {
					break
				}
				return i
			}
			// 還未走完全部節點時如果能源已歸零就再也無法移動
			if sum <= 0 {
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}
