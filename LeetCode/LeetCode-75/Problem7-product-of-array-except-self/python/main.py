import math

class Solution:
    def productExceptSelf(self, nums: list[int]) -> list[int]:
        leftP = [1] * len(nums)
        rightP = [1] * len(nums)
        for i in range(1, len(nums)):
            leftP[i] = nums[i-1] * leftP[i-1]
        for i in range(len(nums)-2, -1, -1):
            rightP[i] = nums[i+1] * rightP[i+1]

        for i in range(len(nums)):
            nums[i] = leftP[i] * rightP[i]
        return nums
    

# Case 1
nums = [1,2,3,4]

# Case 2
# nums = [-1, 1, 0, -3, 3]

solution = Solution()
print(solution.productExceptSelf(nums))
