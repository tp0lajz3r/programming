class Solution:
    def canPlaceFlowers(self, flowerbed: list[int], n: int) -> bool:
        if len(flowerbed) > 1:
            if flowerbed[0] == 0 and flowerbed[1] == 0:
                flowerbed[0] = 1
                n -= 1
            
        for i in range(2, len(flowerbed)-1, 1):
            if flowerbed[i-2] == 0 and flowerbed[i-1] == 0 and flowerbed[i] == 0:
                flowerbed[i-1] = 1
                n -= 1
                
        if flowerbed[len(flowerbed)-2] == 0 and flowerbed[len(flowerbed)-1] == 0:
            flowerbed[len(flowerbed)-1] = 1
            n -= 1

        if n > 0:
            return False
        else:
            return True


# Case 1
# flowerbed = [1,0,0,0,1]
# n = 1

# Case 2
# flowerbed = [1,0,0,0,1]
# n = 2

flowerbed = [0]
n = 1

solution = Solution()
print(solution.canPlaceFlowers(flowerbed, n))