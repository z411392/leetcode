package climbing_stairs

// func factorial(n *big.Int) *big.Int {

// 	if n.Cmp(big.NewInt(1)) <= 0 {
// 		return big.NewInt(1)
// 	}
// 	return big.NewInt(0).Mul(n, factorial(big.NewInt(0).Sub(n, big.NewInt(1))))
// }

// func permutationCount(n *big.Int) *big.Int {
// 	return factorial(n)
// }

// func climbStairs(n int) int {
// 	_2_steps := big.NewInt(int64(n / 2))
// 	_1_steps := big.NewInt(int64(n % 2))
// 	ways := big.NewInt(0)
// 	for {
// 		if _2_steps.Cmp(big.NewInt(0)) <= 0 {
// 			break
// 		}
// 		count := big.NewInt(0).Div(
// 			permutationCount(big.NewInt(0).Add(_2_steps, _1_steps)),
// 			(big.NewInt(0).Mul(permutationCount(_2_steps), permutationCount(_1_steps))),
// 		)
// 		ways = big.NewInt(0).Add(ways, count)
// 		_2_steps = big.NewInt(0).Sub(_2_steps, big.NewInt(1))
// 		_1_steps = big.NewInt(0).Add(_1_steps, big.NewInt(2))
// 	}
// 	return int(ways.Int64()) + 1
// }

/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
*/
var memorized = map[int]int{}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	if value, exists := memorized[n]; exists {
		return value
	}
	memorized[n] = climbStairs(n-1) + climbStairs(n-2)
	return memorized[n]
}
