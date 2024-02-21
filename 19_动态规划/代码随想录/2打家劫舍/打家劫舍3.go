package main

/*
 树形dp
1) 确定递归函数的参数和返回值
这里我们要求一个节点 偷与不偷的两个状态所得到的金钱，那么返回值就是一个长度为2的数组。
其实这里的返回数组就是dp数组。

所以dp数组（dp table）以及下标的含义：下标为0记录不偷该节点所得到的的最大金钱，下标为1记录偷该节点所得到的的最大金钱。

2) 确定终止条件
在遍历的过程中，如果遇到空节点的话，很明显，无论偷还是不偷都是0，所以就返回

3)确定遍历的顺序
首先明确的是使用后序遍历。 因为要通过递归函数的返回值来做下一步计算。
通过递归左节点，得到左节点偷与不偷的金钱。
通过递归右节点，得到右节点偷与不偷的金钱。

4）确定单层递归的逻辑
如果是偷当前节点，那么左右孩子就不能偷，val1 = cur->val + left[0] + right[0]; （如果对下标含义不理解就再回顾一下dp数组的含义）

如果不偷当前节点，那么左右孩子就可以偷，至于到底偷不偷一定是选一个最大的，所以：val2 = max(left[0], left[1]) + max(right[0], right[1]);

最后当前节点的状态就是{val2, val1}; 即：{不偷当前节点得到的最大金钱，偷当前节点得到的最大金钱}
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob_3(root *TreeNode) int {
	res := robTree(root)
	return max(res[0], res[1])
}

func max_3(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func robTree(cur *TreeNode) []int {
	if cur == nil {
		return []int{0, 0}
	}
	left := robTree(cur.Left)
	right := robTree(cur.Right)
	robCur := cur.Val + left[0] + right[0] //这里的递归是从下往上得到的
	notRobCur := max(left[0], left[1]) + max(right[0], right[1])
	return []int{notRobCur, robCur}
}
