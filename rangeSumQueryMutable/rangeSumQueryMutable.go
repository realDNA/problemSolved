"""
Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

The update(i, val) function modifies nums by updating the element at index i to val.

Example:

Given nums = [1, 3, 5]

sumRange(0, 2) -> 9
update(1, 2)
sumRange(0, 2) -> 8
Note:

The array is only modifiable by the update function.
You may assume the number of calls to update and sumRange function is distributed evenly.
"""

class NumArray(object):

    def __init__(self, nums):
        """
        :type nums: List[int]
        """
        len_nums = len(nums)
        indexed_binary_tree = [0]
        indexed_binary_tree[1:len_nums] = nums[:]
        
        for i in range(len_nums+1):
            j = i + (i & (-i))
            if j <= len_nums:              
                indexed_binary_tree[j] = indexed_binary_tree[j] + indexed_binary_tree[i]

        self.indexed_binary_tree = indexed_binary_tree
        self.nums = nums

    def update(self, i, val):
        """
        :type i: int
        :type val: int
        :rtype: void
        """
        
        diff = val - self.nums[i]
        self.nums[i] = val
        i += 1 
        
        while i < len(self.indexed_binary_tree):         
            self.indexed_binary_tree[i] += diff
            i = i + (i & (-i))

    def sumRange(self, i, j):
        """
        :type i: int
        :type j: int
        :rtype: int
        """
        j += 1
        behind = 0
        before = 0

        while j > 0:
            behind += self.indexed_binary_tree[j]
            j = j - (j & (-j))
            
        while i > 0:
            before += self.indexed_binary_tree[i]
            i = i - (i & (-i))
            
        return behind - before
        


# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# obj.update(i,val)
# param_2 = obj.sumRange(i,j)
