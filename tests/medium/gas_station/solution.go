package gas_station

/*
There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].

You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith station to its next (i + 1)th station. You begin the journey with an empty tank at one of the gas stations.

Given two integer arrays gas and cost, return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1. If there exists a solution, it is guaranteed to be unique.
*/
func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, currentGas := 0, 0
	startStation := 0

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i] - cost[i]
		currentGas += gas[i] - cost[i] // 只要確保過程中 currentGas 不出現負數即可
		if currentGas < 0 {            // 一但出現負數就將 startStation 換一個位置並將 currentGas 清空
			startStation = i + 1
			currentGas = 0
		}
	}
	// 如果 totalGas >= 0 則一定存在解
	if totalGas < 0 {
		return -1
	}
	return startStation
}
