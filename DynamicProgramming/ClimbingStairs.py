class Solution:
    def climbStairs(self, n: int) -> int:
        memo = {}
        def f(n):
            if n == 0 or n == 1:
                return 1
            if n not in memo:
                memo[n] = f(n - 1) + f(n - 2)
            return memo[n]
        return f(n)

sol = Solution()
print(sol.climbStairs(4))