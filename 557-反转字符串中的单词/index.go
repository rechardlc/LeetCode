//给定一个字符串 s ，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//
//
//
// 示例 1：
//
//
//输入：s = "Let's take LeetCode contest"
//输出："s'teL ekat edoCteeL tsetnoc"
//
//
// 示例 2:
//
//
//输入： s = "God Ding"
//输出："doG gniD"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 5 * 10⁴
// s 包含可打印的 ASCII 字符。
// s 不包含任何开头或结尾空格。
// s 里 至少 有一个词。
// s 中的所有单词都用一个空格隔开。
//
// Related Topics 双指针 字符串 👍 416 👎 0
package main

import "fmt"

func main() {
	var a = reverseWords("Let's take LeetCode contest")
	fmt.Println(a)
}
func reverseWords(s string) string {
	//length, ret, retLen := len(s), make([]byte, 0), 0
	//for i := 0; i < length; {
	//	start := i
	//	// while(i < lenth && i != " ") {}
	//	for i < length && s[i] != ' ' {
	//		// 当前循环完毕，说明i指针已经移动到空字符串(" ")位置
	//		i++
	//	}
	//	for ; start < i; start++ {
	//		// 当前循环结束，将指针start-i位置的字符串发生颠倒
	//		// -start: 减去初始指针的位置，-1由于当前i指针已经在空字符串位置，空字符串不需要调转位置，+retLen: 将s的index移动到已被添加的调转完成的位置，否则一直从s[0]开始添加
	//		ret = append(ret, s[i-start-1+retLen])
	//	}
	//	if i < length && s[i] == ' ' {
	//		ret = append(ret, ' ')
	//		i++
	//		retLen = len(ret)
	//	}
	//}
	//return string(ret)

	newS := []byte(s)
	i := 0
	j := 0
	for i < len(newS) && j < len(newS) {
		if newS[j] != ' ' && j < len(newS)-1 {
			j++
			continue
		}
		w := j + 1
		if newS[j] == ' ' {
			j--
		}
		for i < j {
			newS[i], newS[j] = newS[j], newS[i]
			i++
			j--
		}
		i, j = w, w
	}
	return string(newS)
}
