class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        #делаем решение как в ЕГЭ
        top_row = [1] * n
        for i in range(m - 1):
            cur_row = [1] * n
            for j in range(1, n):
                cur_row[j] = cur_row[j - 1] + top_row[j]
            top_row = cur_row
        return top_row[-1]