package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 		Tree
//		  3
//		/  \
//	  2      1
//	 / \	/ \
//	6	4  7   8

//--------------------------二叉树遍历递归---------------------------------------------

// 递归
// 前序递归 ：先访问根在访问左右 （同理见其他递归）
//func preorderTraversal(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	fmt.Print(root.Val)           // 根
//	preorderTraversal(root.Left)  // 左
//	preorderTraversal(root.Right) // 右
//}

// 非递归

// 前序
func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	stack := make([]*TreeNode, 0)
	for root != nil && len(stack) != 0 {
		// 先放根和左子节点到result
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		// 从stack里拿右子节点
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

// 中序列
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		// 把左子树放到stack
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 弹出，先取左子节点
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		root = node.Right
	}
	return result
}

// 后序，核心：根节点必须在右节点弹出后才能弹出
func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	stack := make([]*TreeNode, 0)
	// 通过lastVisit标识右子节点是否已经弹出
	var lastVisit *TreeNode
	for root != nil || len(stack) != 0 {
		// 把左子树放到stack
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		// 根节点必须在右子节点弹出后才能弹出
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}

//--------------------------DFS深度遍历---------------------------------------------

// 从上到下
func deepRangeFromUpToDown(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

// 从下到上(分治法)，递归返回结果最后合并，举例：快排
func deepRangeFromLowToUp(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	//分治， 合并结果
	left := deepRangeFromLowToUp(root.Left)
	right := deepRangeFromLowToUp(root.Right)

	return append([]int{root.Val}, append(left, right...)...)
}

//--------------------------DFS广度遍历，层次遍历---------------------------------------------

func levelOrder(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {

			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
		result = append(result, list...)
	}

	return result
}

//--------------------------分治法，模拟---------------------------------------------
/*
	先分别处理局部，再合并结果
	func traversal(root *TreeNode) ResultType  {
		// nil or leaf
		if root == nil {
			// do something and return
		}

		// Divide
		ResultType left = traversal(root.Left)
		ResultType right = traversal(root.Right)

		// Conquer
		ResultType result = Merge from left and right

		return result
	}
*/
