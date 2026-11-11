class Solution:
    def increasingTriplet(self, nums: list[int]) -> bool:
        if len(nums) < 3:
            return False

        k = []

        for i in range(2, len(nums)):
            if nums[i] > nums[i-1] > nums[i-2]:
                return True
            elif nums[i] > nums[i-1]:
                print("check1")
                k.append(nums[i-1])
                k.append(nums[i])
            elif nums[i] > nums[i-2]:
                print("check2")
                k.append(nums[i-2])
                k.append(nums[i])
            elif nums[i-1] > nums[i-2]:
                print("check3")
                k.append(nums[i-2])
                k.append(nums[i-1])

        if len(k) >= 3:
            print(len(k))
            return True

        return False



# Case 1
# nums = [1,2,3,4,5]

# Case 2
# nums = [5,4,3,2,1]

# Case 3
# nums = [2,1,5,0,4,6]

nums = [3,3,3,2,0,1,1]

solution = Solution()
print(solution.increasingTriplet(nums))