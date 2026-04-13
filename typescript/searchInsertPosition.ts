function searchInsert(nums: number[], target: number): number {
  let position = 0
  for (let i = 0; i < nums.length; i++) {
    const num = nums[i]!
    if (num === target) {
      return i
    } else if(num < target) {
      position += 1
    }
  } 
  return position
}

console.log(searchInsert([1,3,5,6], 5))
console.log(searchInsert([1,3,5,6], 2))
console.log(searchInsert([1,3,5,6], 7))