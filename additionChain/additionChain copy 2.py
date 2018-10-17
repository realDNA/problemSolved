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
        self.final=[]

    def checkNumberStatus(self, currentNumber, tempNumber, target, numList, numListLen):
        if tempNumber <= target:
            numList.append(tempNumber)
            print("numList after append = ",numList)
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
                        break
                    else:
                        continue
        return None
   
    def findShortestChainHelper(self, target, numList, minLength):
        print("--------------------------------")
        print("target = ",target)
        print("numList = ",numList)
        numListLen = len(numList)
        currentNumber = numList[numListLen-1]
        print("currentNumber = ",currentNumber)
        self.record.append(numList)
        #print("self.record = ", self.record) 
        if currentNumber == target:
            print("inside return numList = ",numList)
            #self.final.append(numList)
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
                print("length1 1 = ",length1)

        if numListLen >= 2:
            tempNumber = currentNumber + numList2[numListLen-2]
            numList2 = self.checkNumberStatus(currentNumber, tempNumber, target, numList2, numListLen)
            if numList2 is not None:
                length2 = self.findShortestChainHelper(target, numList2, minLength)
                print("length2 1 = ",length2)
 
        minLength = min(minLength, length1, length2)
        print("minLength = ",minLength)
        print("length1 = ",length1)
        print("length2 = ",length2)
        #print("self.final = ",self.final)
        return minLength     
       
 
    def findShortestChain(self, target):
        print("yoyo")
        startNumber = 1
        numList = [startNumber]
        return self.findShortestChainHelper(target, numList, math.inf)



if __name__=="__main__":
    additionChain = AdditionChain()
    additionChain.findShortestChain(27)
