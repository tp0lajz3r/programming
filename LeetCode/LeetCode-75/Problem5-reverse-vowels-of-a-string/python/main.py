class Solution:
    def reverseVowels(self, s: str) -> str:
        vowels = ["a", "e", "i", "o", "u"]
        tmp = []
        for c in s:
            for v in vowels:
                if v == c.lower():
                    tmp.append(c)
        tmp.reverse()
        l = list(s)
        count = 0
        for i in range(0, len(l)):
            for v in vowels:
                if v == l[i].lower():
                    l[i] = tmp[count]
                    count+=1
                    break
                    
        s = ''.join(l)
        return s



# Case 1
# s = "IceCreAm"

# # Case 2
s = "ai"

solution = Solution()
print(solution.reverseVowels(s))