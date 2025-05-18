from typing import List

class Solution:
    def rob(self, nums: List[int]) -> int:
        n = len(nums)
        if n == 1:
            return nums[0]
        ans = [0] * n
        ans[0] = nums[0]
        ans[1] = max(nums[0], nums[1])
        for i in range(2, n):
            ans[i] = max(ans[i - 1], nums[i] + ans[i - 2])

        return max(ans)