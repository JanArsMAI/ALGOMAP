class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        from collections import defaultdict
        l = 0
        max_length = 0
        max_count = 0
        count = defaultdict(int)
        for r in range(len(s)):
            count[s[r]] += 1
            max_count = max(max_count, count[s[r]])
            if (r - l + 1 - max_count) > k:
                count[s[l]] -= 1
                l += 1
            max_length = max(max_length, r - l + 1)

        return max_length

sol = Solution()
print(sol.characterReplacement(s = "ABAB", k = 2))