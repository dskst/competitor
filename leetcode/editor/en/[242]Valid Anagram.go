//Given two strings s and t, return true if t is an anagram of s, and false
//otherwise.
//
// An Anagram is a word or phrase formed by rearranging the letters of a
//different word or phrase, typically using all the original letters exactly once.
//
//
// Example 1:
// Input: s = "anagram", t = "nagaram"
//Output: true
//
// Example 2:
// Input: s = "rat", t = "car"
//Output: false
//
//
// Constraints:
//
//
// 1 <= s.length, t.length <= 5 * 10⁴
// s and t consist of lowercase English letters.
//
//
//
// Follow up: What if the inputs contain Unicode characters? How would you
//adapt your solution to such a case?
//
// Related Topics Hash Table String Sorting 👍 11644 👎 372

package main

// leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	// first try
	//rs := []rune(s)
	//rt := []rune(t)
	//
	//ls := len(rs)
	//lt := len(rt)
	//if ls != lt {
	//	return false
	//}
	//
	//ss := make([]string, 0, ls)
	//st := make([]string, 0, lt)
	//for i := 0; i < ls; i++ {
	//	if i+1 < ls && i+1 < lt {
	//		ss = append(ss, string(rs[i:(i+1)]))
	//		st = append(st, string(rt[i:(i+1)]))
	//	} else {
	//		ss = append(ss, string(rs[i:]))
	//		st = append(st, string(rt[i:]))
	//	}
	//}
	//
	//sort.Strings(ss)
	//sort.Strings(st)
	//if strings.Join(ss, "") == strings.Join(st, "") {
	//	return true
	//}
	//return false

	// example
	if len(s) != len(t) {
		return false
	}

	var freq [26]int

	for idx := 0; idx < len(s); idx++ {
		freq[s[idx]-'a']++
		freq[t[idx]-'a']--
	}

	for idx := 0; idx < len(freq); idx++ {
		if freq[idx] != 0 {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
