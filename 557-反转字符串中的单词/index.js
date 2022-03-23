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
// Related Topics 双指针 字符串 👍 417 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * @param {string} s
 * @return {string}
 */
const reverseWords = function(s) {
  let len = s.length, ret = "", retLen = 0
  for (let i = 0; i < len;) {
    let start = i
    while (i < len && s[i] !== " ") {
      i ++
    }
    for (;start < i; start ++) {
      ret += s[i - start - 1 + retLen]
    }
    if (i < len && s[i] === " ") {
      i ++
      ret += " "
      retLen = ret.length
    }
  }
  return ret
};
//leetcode submit region end(Prohibit modification and deletion)
// let a = reverseWordsReplace("Let's {take <LeetCode contest)")
// console.log(a)
function reverseWordsReplace(s) {
  return s.replace(/\S+/ig, match => {
    return match.split("").reverse().join("")
  })
}

const doublePointer = function (s) {
  let i = 0, j = 0, sArr = s.split("")
  while (i < sArr.length && j < sArr.length ) {
    if (sArr[i] !== " " && j < sArr.length - 1) {
      j ++
    } else {
      let w = j + 1 // 当前的j在空字符串位置，w记录下一个非空字符串位置
      if (sArr[i] === " ") {
        j --
      }
      while (i < j) {
        [sArr[j], sArr[i]] = [sArr[i], sArr[j]]
        i ++
        j --
      }
      i = w
      j = w
    }
  }
  return sArr.join("")
}
console.log(doublePointer("Let's {take <LeetCode contest)"))
