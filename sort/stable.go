package sort

/*
	稳定排序：
		冒泡
		插入
		归并
		桶
		基数
*/

// 手里一把扑克牌，第一轮，把最大的放在右边；第二轮，排除最右边一张，接着从0开始找最大，放在最右边...
func Bubble(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 手里一把扑克牌，从左手边第二张（0是第一张）开始，向左比较；然后从左手边第三张开始.....
func Insert(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}
