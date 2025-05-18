class Solution:
    def fib(self, n: int) -> int:
        memoiz_dict = {0:0, 1:1}
        def f(x):
            if x in memoiz_dict:
                return memoiz_dict[x]
            else:
                memoiz_dict[x] = f(x - 1) + f(x - 2)
                return memoiz_dict[x]
        return f(n)
    
sol = Solution()
print(sol.fib(4))