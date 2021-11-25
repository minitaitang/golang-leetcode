https://leetcode-solution-leetcode-pp.gitbook.io/leetcode-solution/medium/bus-fare

You are given a list of sorted integers days , where you must take the bus for on each day. Return the lowest cost it takes to travel for all the days.
There are 3 types of bus tickets.
1 day pass for 2 dollars
7 day pass for 7 dollars
30 day pass for 25 dollars
Constraints
n ≤ 100,000 where n is the length of days
Example 1
Input
days = [1, 3, 4, 5, 29]
Output
9
Explanation
The lowest cost can be achieved by purchasing a 7 day pass in the beginning and then a 1 day pass on the 29th day.
Example 2
Input
days = [1]
Output
2