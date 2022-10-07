package main

import (
	"fmt"
	"github.com/alg/00.lbld/basicStruct"
)

func main() {

	// 01.删除有序数组中的重复项
	//removeDuplicatesTest()

	// 02.删除排序链表中的重复元素
	//deleteDuplicatesTest()

	// 03.移除数组中值为target的所有元素，并返回移除后新数组的长度
	//removeElementTest()

	// 04.将数组中所有值为0的元素移到数组末尾
	//moveZeroesTest()

	// 05. 二分查找
	//binarySearchTest()

	// 06. 两数之和: 数组已经按非递减顺序排列，从数组中找到满足相加之和等于target的两个数的下标，以 [a,b]的形式返回
	//twoSumTest()

	// 07. 反转数组
	//reverseStringTest()

	// 08. 找出字符串中最长的回文串
	longestPalindromeTest()
}

func removeDuplicatesTest() {
	nums := []int{1, 2, 2, 3, 3, 4, 5, 5, 6}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {

	if len(nums) <= 0 {
		return len(nums)
	}

	var slow = 0
	var fast = 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

func deleteDuplicatesTest() {

	node1 := &basicStruct.ListNode{Val: 1}
	node2 := &basicStruct.ListNode{Val: 2}
	node3 := &basicStruct.ListNode{Val: 3}
	node4 := &basicStruct.ListNode{Val: 3}
	node5 := &basicStruct.ListNode{Val: 4}
	node6 := &basicStruct.ListNode{Val: 4}
	node7 := &basicStruct.ListNode{Val: 7}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	newHead := deleteDuplicates(node1)
	newHead.Traverse()
}

func deleteDuplicates(head *basicStruct.ListNode) *basicStruct.ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var slow, fast = head, head
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = fast
		}
		fast = fast.Next
	}

	slow.Next = nil

	return head
}

func removeElementTest() {
	nums := []int{1, 2, 2, 3, 3, 4, 5, 5, 6}
	fmt.Println(removeElement(nums, 5))
	fmt.Println(nums)
}

func removeElement(nums []int, target int) int {

	var slow, fast int
	for fast < len(nums) {
		if nums[fast] != target {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func moveZeroesTest() {
	nums2 := []int{0, 2, 2, 3, 3, 0, 4, 5, 5, 6}
	moveZeroes01(nums2)
	fmt.Println(nums2)

	nums3 := []int{0, 2, 2, 3, 3, 0, 4, 5, 5, 6}
	moveZeroes02(nums3)
	fmt.Println(nums3)

	nums4 := []int{0, 2, 2, 3, 3, 0, 4, 5, 5, 6}
	moveZeroes03(nums4)
	fmt.Println(nums4)
}

// 将所有非零元素挪到前面来，并将原来的位置置为0即可，保证非零顺序
func moveZeroes01(nums []int) {

	var slow, fast int
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			if slow != fast {
				nums[fast] = 0
			}
			slow++
		}
		fast++
	}
	return
}

// 将所有零元素移到末尾，两种方案优劣对比需要看数组中0元素的占比，超过一半就moveZeroes01更优，小于一半就moveZeroes02更优
// 但是moveZeroes02改变了非0元素的顺序，这可能会有问题
func moveZeroes02(nums []int) {

	var slow, fast = len(nums) - 1, len(nums) - 1
	for fast >= 0 {
		if nums[fast] == 0 {
			basicStruct.SwapInt(nums, slow, fast)
			slow--
		}
		fast--
	}
	return
}

// 移除数组中所有值为0的元素，并将末尾的值置位0，保证非0顺序
func moveZeroes03(nums []int) {
	newLen := removeElement(nums, 0)

	for i := newLen; i < len(nums); i++ {
		nums[i] = 0
	}
	return
}

func binarySearchTest() {

	nums := []int{0, 2, 2, 3, 3, 0, 4, 5, 5, 6}
	fmt.Println(binarySearch(nums, 82))
}

func binarySearch(nums []int, target int) int {

	var left int
	var right = len(nums) - 1

	for left <= right {

		var mid = left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	return -1
}

func twoSumTest() {

	nums := []int{0, 2, 2, 3, 3, 0, 4, 5, 5, 6}
	fmt.Println(twoSum(nums, 9))
}

func twoSum(nums []int, target int) []int {
	var left = 0
	var right = len(nums) - 1

	for left < right { // 不能相等，相等就是同一个元素重复使用了
		if nums[left]+nums[right] == target {
			return []int{left, right}
		} else if nums[left]+nums[right] > target {
			right--
		} else if nums[left]+nums[right] < target {
			left++
		}
	}
	return []int{-1, -1}
}

func reverseStringTest() {
	s := "hello"
	fmt.Println(string(reverseString([]byte(s))))
}

func reverseString(s []byte) []byte {
	var left int
	var right = len(s) - 1

	for left < right {
		basicStruct.SwapBytes(s, left, right)
		left++
		right--
	}
	return s
}

func longestPalindromeTest() {

	var s = "helleh"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {

	var res string

	for i := 0; i < len(s); i++ {
		res1 := getLongestPalindrome(s, i, i)		// 字符串为奇数
		res2 := getLongestPalindrome(s, i, i+1)		// 字符串为偶数

		if len(res1) > len(res2) {
			if len(res) < len(res1) {
				res = res1
			}
		} else {
			if len(res) < len(res2) {
				res = res2
			}
		}
	}
	return res
}

// 找出以left,right为中心的最长回文串
func getLongestPalindrome(s string, left, right int) string {

	var res string
	for left >= 0 && right < len(s) && s[left] == s[right] {
		res = s[left : right+1]
		left--
		right++
	}
	return res
}
