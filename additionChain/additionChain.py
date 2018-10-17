'''
Efficient exponentiation
Problem 122 
The most naive way of computing n15 requires fourteen multiplications:

n × n × ... × n = n15

But using a "binary" method you can compute it in six multiplications:

n × n = n2
n2 × n2 = n4
n4 × n4 = n8
n8 × n4 = n12
n12 × n2 = n14
n14 × n = n15

However it is yet possible to compute it in only five multiplications:

n × n = n2
n2 × n = n3
n3 × n3 = n6
n6 × n6 = n12
n12 × n3 = n15

We shall define m(k) to be the minimum number of multiplications to compute nk; for example m(15) = 5.

For 1 ≤ k ≤ 200, find ∑ m(k).
'''

import math 

class AdditionChain(object):
    def __init__(self):
        self.record=[]

    def checkNumberStatus(self, currentNumber, tempNumber, target, numList, numListLen):
        if tempNumber <= target:
            numList.append(tempNumber)
            if numList not in self.record:
                return numList
        else:
            if numListLen >= 3:
                for n in range(numListLen-2,-1,-1):
                    tempNumber = currentNumber + numList[n]
                    if tempNumber <= target:
                        numList.append(tempNumber)
                        if numList in self.record:
                            break
                        return numList
                    else:
                        continue
        return None
   
    def findShortestChainHelper(self, target, numList, minLength):
        numListLen = len(numList)
        currentNumber = numList[numListLen-1]
        self.record.append(numList)
        
        if currentNumber == target:
            return numListLen
      
        length1 = math.inf
        length2 = math.inf
        numList1 = numList[:]
        numList2 = numList[:]
 
        tempNumber = currentNumber*2
        if tempNumber <= target:
            numList1 = self.checkNumberStatus(currentNumber, tempNumber, target, numList1, numListLen)
            if numList1 is not None: 
                length1 = self.findShortestChainHelper(target, numList1, minLength)

        if numListLen >= 2:
            tempNumber = currentNumber + numList2[numListLen-2]
            numList2 = self.checkNumberStatus(currentNumber, tempNumber, target, numList2, numListLen)
            if numList2 is not None:
                length2 = self.findShortestChainHelper(target, numList2, minLength)
        
        minLength = min(minLength, length1, length2)
        return minLength     
       
 
    def findShortestChain(self, target):
        startNumber = 1
        numList = [startNumber]
        self.record = []
        if target == 1:
            return 1 # k=1
        return self.findShortestChainHelper(target, numList, math.inf) - 1 # -1 means don't count k=1



if __name__=="__main__":
    additionChain = AdditionChain()
    sumResult = 0
    for i in range(1,201):
        minLength = additionChain.findShortestChain(i)
        sumResult += minLength
        print(f"number {i}: {minLength}")
    print("final sum is ", sumResult)
