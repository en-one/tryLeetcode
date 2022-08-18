package sort

/*
	不稳定排序：
		选择
		希尔
		快排序
		堆
*/

// 手里一把扑克牌，把最小的放在最左边，然后从下一张开始找最小的
func Select(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

// 加强版分组插入，按照3x+1开始递增排序
func Shell(arr []int) []int {
	h := 1
	for h < len(arr)/3 {
		h = h*3 + 1
	}

	for h >= 1 {
		// 插入排序
		for i := h; i < len(arr); i++ {
			for j := i; j > h-1; j -= h {
				if arr[j] < arr[j-h] {
					arr[j], arr[j-h] = arr[j-h], arr[j]
				}
			}
		}
		h = h / 3
	}
	return arr
}

// 以第一个为基准，大于第一个放右边，小于第一个放左边，分别对左右递归，然后集合
func Quick(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for i := 1; i < len(arr); i++ {
		if arr[i] > pivot {
			right = append(right, arr[i])
		} else {
			left = append(left, arr[i])
		}
	}

	left = Quick(left)
	right = Quick(right)

	return append(left, append([]int{pivot}, right...)...)
}

func Heap(arr []int) []int {
	len := len(arr)
	last_node := len - 1
	last_node_parent := last_node / 2

	// 构建二叉堆
	for i := last_node_parent; i >= 0; i-- {
		heapy(arr, len, i)
	}

	// 调换最后一个与堆顶位置，然后重排堆
	for i := last_node; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapy(arr, i, 0)
	}
	return arr
}

func heapy(tree []int, num, index int) []int {
	if index >= num {
		return tree
	}

	lchild := 2*index + 1
	rchild := 2*index + 2
	max := index

	if lchild < num && tree[lchild] > tree[max] {
		max = lchild
	}
	if rchild < num && tree[rchild] > tree[max] {
		max = rchild
	}
	if max != index {
		tree[max], tree[index] = tree[index], tree[max]
		heapy(tree, num, max)
	}
	return tree
}
