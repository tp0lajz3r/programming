class Solution:
    def kidsWithCandies(self, candies: list[int], extraCandies: int) -> list[bool]:
        tmp = sorted(candies, reverse=True)
        max = tmp[0] - extraCandies
        res = []
        for candy in candies:
            if candy < max:
                res.append(False)
            else:
                res.append(True)
        return res

# Case 1
# candies = [2,3,5,1,3]
# extraCandies = 3

# Case 2
# candies = [4,2,1,1,2]
# extraCandies = 1

# Case 3
candies = [12,1,12]
extraCandies = 10

solution = Solution()
print(solution.kidsWithCandies(candies, extraCandies))