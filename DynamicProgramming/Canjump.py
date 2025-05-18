class Solution(object):
    def canJump(self, nums):
        n = len(nums)
        farthest = 0
        for i in range(n):
            if i > farthest:
                return False
            farthest = max(farthest, i + nums[i])
        return True