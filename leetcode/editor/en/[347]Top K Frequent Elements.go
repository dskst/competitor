//Given an integer array nums and an integer k, return the k most frequent
//elements. You may return the answer in any order.
//
//
// Example 1:
// Input: nums = [1,1,1,2,2,3], k = 2
//Output: [1,2]
//
// Example 2:
// Input: nums = [1], k = 1
//Output: [1]
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
// k is in the range [1, the number of unique elements in the array].
// It is guaranteed that the answer is unique.
//
//
//
// Follow up: Your algorithm's time complexity must be better than O(n log n),
//where n is the array's size.
//
// Related Topics Array Hash Table Divide and Conquer Sorting Heap (Priority
//Queue) Bucket Sort Counting Quickselect 👍 16788 👎 622

package main

// leetcode submit region begin(Prohibit modification and deletion)
func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	cnt := make(map[int][]int, len(nums)+1)
	var r []int

	for _, num := range nums {
		freq[num]++
	}

	for k, v := range freq {
		cnt[v] = append(cnt[v], k)
	}

	for i := len(nums); i > 0; i-- {
		r = append(r, cnt[i]...)
		if len(r) == k {
			break
		}
	}

	return r
}

//leetcode submit region end(Prohibit modification and deletion)
