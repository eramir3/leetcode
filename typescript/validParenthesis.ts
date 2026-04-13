function isValidParenthesis(s: string): boolean {
 const pairs: Record<string, string> = {
   ")": "(",
   "}": "{",
   "]": "["
 }

 const stack: string[] = []
 for (const char of s) {
   if (char in pairs) {
     // closing bracket → check match
     const expected = pairs[char]
     const top = stack.pop();
     if (top !== expected) {
       return false;
     }
   } else {
     // opening bracket → push to stack
     stack.push(char);
   }
 }
 return stack.length === 0
}

console.log(isValidParenthesis("({[]})"))
