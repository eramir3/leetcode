function romanToInt(s: string): number {
  const map: Record<string, number> = {
    I: 1,
    V: 5,
    X: 10,
    L: 50,
    C: 100,
    D: 500,
    M: 1000,
  }

  let sum = 0
  let last = 0
  for (let i = s.length - 1; i >= 0; i--) {
    const current = map[s[i]!]!
    if (current >= last) {
      sum += current
    } else {
      sum -= current
    }
    last = current
  }
  return sum
}

//romanToInt("I")
console.log(romanToInt("XCVIII"))