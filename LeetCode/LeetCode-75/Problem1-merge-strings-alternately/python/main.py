class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        res = ""
        if len(word1) >= len(word2):
            for i in range(0, len(word2)):
                res += word1[i] + word2[i]
            res += word1[len(word2):]
        else:
            for i in range(0, len(word1)):
                res += word1[i] + word2[i]
            res += word2[len(word1):]
        return res


# Case 1
word1 = "abc"
word2 = "pqr"

# Case 2
# word1 = "ab"
# word2 = "pqrs"

# Case 3
# word1 = "abcd"
# word2 = "pq"

solution = Solution()
print(solution.mergeAlternately(word1, word2))