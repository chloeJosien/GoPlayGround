// problem found at https://leetcode.com/problems/length-of-last-word/

package leetCodeLenLastWord

import "strings"

func lengthOfLastWord(s string) int {
	var trimEndings = strings.TrimSpace(s)
	var listOfStrings = strings.Split(trimEndings, " ")
	var lastWord = listOfStrings[len(listOfStrings)-1]

	return len(lastWord)
}
