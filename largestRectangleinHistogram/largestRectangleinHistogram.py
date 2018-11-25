"""
Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.

 


Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].

 


The largest rectangle is shown in the shaded area, which has area = 10 unit.

 

Example:

Input: [2,1,5,6,2,3]
Output: 10	
"""

class Solution:
    def largestRectangleArea(self, heights):
        """
        :type heights: List[int]
        :rtype: int
        """
        if not heights:
            return 0
        
        stack = [-1]
        heights.append(0)
        heights_len = len(heights) 
        max_area = 0
        
        for i in range(heights_len):
            while stack and heights[i] < heights[stack[-1]]:
                h = heights[stack.pop()]
                w = i - stack[-1] - 1
                area = h*w
                max_area = max(max_area, area)
            stack.append(i)
        return max_area

if __name__=="__main__":
    heights = [2,1,4,5,6]
    s = Solution()
    result = s.largestRectangleArea(heights)
    print("result = ", result)
