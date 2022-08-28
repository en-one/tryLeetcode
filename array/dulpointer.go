package array

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

// -------------------- 左右指针----------------------------
