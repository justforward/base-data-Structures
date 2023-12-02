package main

import (
	"math/rand"
)

// 0(1)时间复杂度内完成 查询、删除、插入
// 1、哈希记录加入和删除的数，可以在O（1）的时间复杂度内完成
// 2、使用数组维护所有的数，方便随机取一个数，数组后加入一个数也是O（1）
// 3、删除的时候，如果删除的值不是最后一个，那么得倒的结果与最后一个值进行交换，删除的数是数组的最后一个，可以O（1删除）
type RandomizedSet struct {
	tmpMap map[int]int
	nums   []int
}

func Constructor() RandomizedSet {

	return RandomizedSet{make(map[int]int), make([]int, 0)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.tmpMap[val]; ok {
		return false
	}
	this.tmpMap[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.tmpMap[val]
	if !ok {
		return false
	}
	last := len(this.nums) - 1
	this.nums[idx] = this.nums[last]
	this.tmpMap[this.nums[idx]] = idx
	this.nums = this.nums[:last]
	delete(this.tmpMap, val)
	return true

}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
