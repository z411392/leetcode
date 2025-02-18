package insert_delete_getrandom_o1

import (
	"math/rand"
	"slices"
	"time"
)

type RandomizedSet struct {
	Values map[int]int
	Keys   []int
	Random *rand.Rand
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Values: make(map[int]int),
		Keys:   []int{},
		Random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, exists := this.Values[val]
	if exists {
		return false
	}
	this.Values[val] = val
	this.Keys = append(this.Keys, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	_, exists := this.Values[val]
	if !exists {
		return false
	}
	index := slices.Index(this.Keys, val)
	this.Keys = append(this.Keys[:index], this.Keys[index+1:]...)
	delete(this.Values, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	val := this.Keys[this.Random.Intn(len(this.Keys))]
	return val
}
