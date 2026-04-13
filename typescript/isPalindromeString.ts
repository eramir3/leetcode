function isPalindromeString(text: string) {
  const string1 = text.replace(/[^a-zA-Z0-9]/g, "").toLowerCase()
  const string2 = text.replace(/[^a-zA-Z0-9]/g, "").toLowerCase().split("").reverse().join("")
  return string1 == string2
}

function isPalindromeString2(text: string) {
  const cleaned = text.replace(/[^a-zA-Z0-9]/g, "").toLowerCase()
  const chars: string[] = Array.from(cleaned)
  for (let i = 0, j = cleaned.length - 1; i < j; i++, j--) {
    [chars[i], chars[j]] = [chars[j]!, chars[i]!]
  }
  return cleaned == chars.join("")
}

console.log(isPalindromeString("A man, a plan, a canal - Panama!"))
console.log(isPalindromeString2("A man, a plan, a canal - Panama!"))
