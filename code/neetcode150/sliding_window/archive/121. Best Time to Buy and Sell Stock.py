from typing import List

class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        profit = 0
        buy = (0,prices[0])
        sell = (0,0)

        for day,price in enumerate(prices):
            if price < buy[1]:
                buy = (day,price)
            
            if price - buy[1] > profit and day != buy[0]:
                profit = price - buy[1]
                sell = (day,price)
        
        return profit
    
    def betterMaxProfit(self, prices: List[int]) -> int:
        # use left/right pointers to compare profit
        buy,sell = 0,1
        maxP = 0

        # iterate through array
        while sell < len(prices):
            # check to make sure profit will be positive
            if prices[buy] < prices[sell]:
                profit = prices[sell] - prices[buy]
                maxP = max(profit, maxP)
            else:
            # if its negative - move buy pointer
                buy = sell
            
            sell += 1

print(Solution().maxProfit([1,2]))