function lengthOfLastWord(s: string): number {
  const text = s.trim().split(" ")
  return text[text.length-1]!.length
}

console.log(lengthOfLastWord("   fly me   to   the moon  "))