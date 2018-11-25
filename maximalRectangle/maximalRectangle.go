"""
Given a 2D binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's and return its area.

Example:

Input:
[
  ["1","0","1","0","0"],
  ["1","0","1","1","1"],
  ["1","1","1","1","1"],
  ["1","0","0","1","0"]
]
Output: 6
"""

class Solution:
    def largestRectangleArea(self, heights, heights_len):
        """
        :type heights: List[int]
        :rtype: int
        """
        if not heights:
            return 0
        
        stack = [-1]
        max_area = 0
        
        for i in range(heights_len):
            while heights[i] < heights[stack[-1]]:
                h = heights[stack.pop()]
                w = i - stack[-1] - 1
                area = h*w
                max_area = max(max_area, area)
            stack.append(i)
        
        return max_area

    def maximalRectangle(self, matrix):
        """
        :type matrix: List[List[str]]
        :rtype: int
        """
        if not matrix:
            return 0
        
        max_rec = 0 
        row_len = len(matrix)
        col_len = len(matrix[0])  
        heights = [0]*(col_len+1)
        heights_len = col_len+1
         
        for row in matrix:
            for col in range(col_len):
                heights[col] = heights[col] + 1 if row[col] == '1' else 0
            rec = self.largestRectangleArea(heights, heights_len)
            max_rec = max(max_rec,rec)
            
        return max_rec
            
        
