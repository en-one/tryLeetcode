package array

import "fmt"

// 索引作为指针

// -------------------- 快慢指针，原地修改数组----------------------------

// 1、删除重复元素,返回数组长度，保留重复元素的初始值 1,1,2,2,2,3
func removeDuplicates(nums []int) int {
	len := len(nums)
	if len == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}

	return slow + 1
}

// 2、移除指定元素，返回长度  nums = [0,1,2,2,3,0,4,2], val = 2
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// 3、所有 0 移动到数组的末尾 : nums = [0,1,0,3,12]  [1,3,12,0,0]
func moveZeroes(nums []int) {
	i := removeElement(nums, 0)
	for ; i < len(nums); i++ {
		nums[i] = 0
	}
	fmt.Println(nums)
}

// -------------------- 左右指针----------------------------
