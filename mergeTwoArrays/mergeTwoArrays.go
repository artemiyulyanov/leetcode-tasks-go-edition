package mergeTwoArrays

type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrayToLinkedNode(array []int) *ListNode {
	var headArray *ListNode = &ListNode{}
	var curArray = headArray

	for _, value := range array {
		curArray.Val = value
		curArray.Next = &ListNode{}

		curArray = curArray.Next
	}

	curArray.Next = nil

	return headArray
}

//func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//
//	return &headMergedList
//}
