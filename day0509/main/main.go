package main

import (
	. "Code_Random_Thoughts/mypkg"
	"math"
)

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
func getMinDfs(pre, cur **TreeNode, min *int) {
	if (*cur).Left != nil {
		getMinDfs(pre, &(*cur).Left, min)
	}
	if (*pre) != nil {
		sub := abs((*pre).Val - (*cur).Val)
		if sub < (*min) {
			*min = sub
		}
	}
	*pre = *cur
	if (*cur).Right != nil {
		getMinDfs(pre, &(*cur).Right, min)
	}
}

func getMinimumDifference(root *TreeNode) int {
	var pre *TreeNode
	res := math.MaxInt
	getMinDfs(&pre, &root, &res)
	return res
}

func findMode(root *TreeNode) []int {
	modeS := make([]int, 0)
	modeCount := 0
	curNum := 0
	curCount := 0
	var pre *TreeNode
	var dfsMode func(node *TreeNode)
	dfsMode = func(node *TreeNode) {
		if node.Left != nil {
			dfsMode(node.Left)
		}
		//fmt.Println(node.Val)
		if pre == nil {
			modeS = []int{node.Val}
			modeCount = curCount
			curNum = node.Val
			curCount = 1
		} else {
			if pre.Val == node.Val {
				curCount++
			} else {
				if curCount > modeCount {
					modeS = []int{curNum}
					modeCount = curCount
					curNum = node.Val
					curCount = 1
				} else if curCount == modeCount {
					modeS = append(modeS, curNum)
					curNum = node.Val
					curCount = 1
				} else {
					curNum = node.Val
					curCount = 1
				}
			}
		}
		pre = node
		if node.Right != nil {
			dfsMode(node.Right)
		}
	}
	dfsMode(root)
	if modeS[len(modeS)-1] != curNum {
		if curCount > modeCount {
			modeS = []int{curNum}
		} else if curCount == modeCount {
			modeS = append(modeS, curNum)
		}
	}
	return modeS
}

/*
find p = 1
find q = 2
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var res *TreeNode
	var dfsCommonNode func(root, p, q *TreeNode) int
	dfsCommonNode = func(root, p, q *TreeNode) int {
		if root == nil {
			return -1
		}
		if root == p {
			left := dfsCommonNode(root.Left, p, q)
			if left == 2 {
				res = root
				return 0
			}
			right := dfsCommonNode(root.Right, p, q)
			if right == 2 {
				res = root
				return 0
			}
			return 1
		} else if root == q {
			left := dfsCommonNode(root.Left, p, q)
			if left == 1 {
				res = root
				return 0
			}
			right := dfsCommonNode(root.Right, p, q)
			if right == 1 {
				res = root
				return 0
			}
			return 2
		} else {
			left := dfsCommonNode(root.Left, p, q)
			if left == 0 {
				return 0
			}
			right := dfsCommonNode(root.Right, p, q)
			if right == 0 {
				return 0
			}
			if (left + right) == 3 {
				res = root
				return 0
			}
			if left == 1 || right == 1 {
				return 1
			} else if left == 2 || right == 2 {
				return 2
			} else {
				return -1
			}
		}
	}
	dfsCommonNode(root, p, q)
	return res
}
