import re

class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = re.sub(r'[^A-Za-z0-9]+', '', s.strip().lower())

        if len(s) == 0:
            return True

        i = 0
        j = len(s) - 1
        while j - i > 0:
            if (s[i] != s[j]):
                return False
            i+=1
            j-=1
            
        return True



sol = Solution()
print(sol.isPalindrome("aa"))