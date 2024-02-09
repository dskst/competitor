// Given an integer array nums, return true if any value appears at least twice
// in the array, and return false if every element is distinct.
//
// Example 1:
// Input: nums = [1,2,3,1]
// Output: true
//
// Example 2:
// Input: nums = [1,2,3,4]
// Output: false
//
// Example 3:
// Input: nums = [1,1,1,3,3,4,3,2,4,2]
// Output: true
//
// Constraints:
//
// 1 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
//
// Related Topics Array Hash Table Sorting 👍 11487 👎 1271
package main

// leetcode submit region begin(Prohibit modification and deletion)
func containsDuplicate(nums []int) bool {
	// first try
	//m := make(map[int]bool)
	//var uniq []int
	//
	//for _, v := range nums {
	//	if !m[v] {
	//		m[v] = true
	//		uniq = append(uniq, v)
	//	} else {
	//		return true
	//	}
	//}
	//return len(nums) != len(uniq)

	// example
	if len(nums) <= 1 {
		return false
	}

	xm := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := xm[v]; ok {
			return true
		}

		xm[v] = struct{}{}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
