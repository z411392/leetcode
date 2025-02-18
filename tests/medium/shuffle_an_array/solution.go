package shuffle_an_array

import (
	"math/rand"
	"time"
)

type Solution struct {
	nums   []int
	Random *rand.Rand
}

func Constructor(nums []int) Solution {
	return Solution{
		nums:   nums,
		Random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	indexes := this.Random.Perm(len(this.nums))
	values := make([]int, len(this.nums))
	for i, index := range indexes {
		values[i] = this.nums[index]
	}
	return values
}
