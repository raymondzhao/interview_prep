class Solution:

    @staticmethod
    def isValid(s: str) -> bool:
        stack = []
        for i in range(len(s)):
            if s[i] in "([{":
                stack.append(s[i])
                
            elif s[i] == ")" and len(stack) > 0 and stack[-1] == "(":
                stack.pop()
            
            elif s[i] == "]" and len(stack) > 0 and stack[-1] == "[":
                stack.pop()
            
            elif s[i] == "}" and len(stack) > 0 and stack[-1] == "{":
                stack.pop()

            else:
                return False

        return len(stack) == 0

# print(Solution.isValid("()"))
# print(Solution.isValid("()[]{}"))
# print(Solution.isValid("([)]"))
print(Solution.isValid("]"))