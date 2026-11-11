from math import gcd

class Solution:
    def gcdOfStrings(self, str1: str, str2: str) -> str:
        if str1 + str2 != str2 + str1:
            return ""
        else:
            length = gcd(len(str1), len(str2))
            return str1[:length]       


# Case 1
str1 = "ABCABC"
str2 = "ABC"

# Case 2
# str1 = "ABABAB"
# str2 = "ABAB"

# Case 3
# str1 = "LEET"
# str2 = "CODE"

# Case 4
# str1 = "AAAAAB"
# str2 = "AAA"

solution = Solution()
print(solution.gcdOfStrings(str1, str2))