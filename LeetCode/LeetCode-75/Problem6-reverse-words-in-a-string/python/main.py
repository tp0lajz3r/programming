class Solution:
    def reverseWords(self, s: str) -> str:
        s = s.split()
        s.reverse()
        s = " ".join(s)
        return s


# Case 1
s = "the sky is blue"

# Case 2
# s = "  hello world  "

# Case 3
# s = "a good   example"


solution = Solution()
print(solution.reverseWords(s))