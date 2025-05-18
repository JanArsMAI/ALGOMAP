class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        all_set = set()
        lobgest = 0
        l = 0
        for right in range(len(s)):
            while s[right] in all_set:
                all_set.remove(s[l])
                l += 1
            longest = max(lobgest, (right - l) + 1)
            all_set.add(s[right])
        return lobgest


sol = Solution()
print(sol.longestOnes(nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2))