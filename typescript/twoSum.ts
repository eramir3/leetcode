function twoSum(nums: number[], target: number): number[] | null {
  const map = new Map<number, number>()
  for (let i = 0; i < nums.length; i++) {
      const num = nums[i]
      const complement = target - num!
      if (map.has(complement)) {
          return [map.get(complement)!, i]
      }
      map.set(num!, i)
  }
  return null
}
console.log(twoSum([2, 7, 11, 15], 9));