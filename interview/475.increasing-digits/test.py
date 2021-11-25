def solve(n, k):
    def dp(i, cnt):
        if i == n:
            return 1
        ans = dp(i + 1, 0)  # place L
        if cnt < k:
            ans += dp(i + 1, cnt + 1)  # place W if I can
        return ans
    return dp(0, 0) % (10 ** 9 + 7)

print(solve(3,2))