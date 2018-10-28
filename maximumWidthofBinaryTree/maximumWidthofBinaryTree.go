/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func compareRange(mostLeft int, position int, maxRange int) int {
    curRange := position - mostLeft
    if curRange > maxRange {
        maxRange = curRange
    }
    return maxRange
    
}

func widthOfBinaryTree(root *TreeNode) int {
    var treeNodes []*TreeNode
    var depths []uint
    var depth uint  
    var positions []int
    
    treeNodes = append(treeNodes, root)
    depth = 1
    depths = append(depths, depth)
    position := 0
    positions = append(positions, position)
    mostLeft := 0  
    maxRange := 0
    treeNodesNum := len(treeNodes)
    
    for i:=0; i<treeNodesNum; i++ {
        treeNode := treeNodes[i]
        curDepth := depths[i]
        curposition := positions[i]

        if curDepth != depth {
            maxRange = compareRange(mostLeft, position, maxRange)
            depth++
            mostLeft = curposition
        }
        position = curposition

        if treeNode != nil {
            treeLeft := treeNode.Left
            if treeLeft != nil {
                treeNodes = append(treeNodes, treeLeft)
                depths = append(depths, depth+1)
                position1 := 2*position
                positions = append(positions, position1)  
            }
            
            treeRight := treeNode.Right
            if treeRight != nil {
                treeNodes = append(treeNodes, treeRight)
                depths = append(depths, depth+1)
                position2 := 2*position + 1
                positions = append(positions, position2)  
            }
            treeNodesNum = len(treeNodes)
        }
    }
    maxRange = compareRange(mostLeft, position, maxRange)
    return maxRange + 1
}