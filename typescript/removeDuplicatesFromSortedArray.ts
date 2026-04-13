function removeDuplicates(nums: number[]): number {
  let k = 0
  let pointer = undefined
  for (const num of nums) {
    if (num != pointer) {
      nums[k] = num!
      pointer = num!
      k++
    }
  }
  return k
}

console.log(removeDuplicates([0,0,1,1,1,2,2,3,3,4]))
console.log(removeDuplicates([1,1,2]))