package test

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"
)

func TestJson(t *testing.T) {
	s := []string{"abdd", "shsjhs", "hdhiuhjdiu", "djuhduir"}
	sjson, _ := json.Marshal(s)
	fmt.Println(string(sjson))
	var s1 []string
	if err := json.Unmarshal(sjson, &s1); err != nil {
		fmt.Println(err)
	}
	fmt.Println(s1)
}

func TestString(t *testing.T) {
	//buildTree([]int{-1}, []int{-1})
	fmt.Println(math.MaxInt32)
}

var cur int = 0

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[cur], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[cur] {
			break
		}
	}
	// fmt.Printf("pre:%v  inorder:%v\n",preorder,inorder)
	cur++
	root.Left = buildTree(preorder, inorder[0:i])
	if i < len(inorder) {
		root.Right = buildTree(preorder, inorder[i+1:])
	}
	return root
}
