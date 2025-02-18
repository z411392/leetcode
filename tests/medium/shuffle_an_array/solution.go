package shuffle_an_array

import (
	"iter"
	"math/rand"
	"slices"
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
	iter := iter.Seq[int](func(yield func(num int) bool) {
		indexes := this.Random.Perm(len(this.nums))
		// fmt.Printf("%v\n", indexes)
		for _, index := range indexes {
			keepGoing := yield(this.nums[index])
			if !keepGoing {
				return
			}
		}
	})
	return slices.AppendSeq([]int{}, iter)
}
