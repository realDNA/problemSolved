import math 

class AdditionChain(object):
    def __init__(self):
        self.record=[]
        self.final=[]

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
            numList1.append(tempNumber)
            if numList1 not in self.record:
                
                length1 = self.findShortestChainHelper(target, numList1, minLength)
                print("length1 1 = ",length1)
        else:
            if numListLen >= 3:
                for n in range(numListLen-2,-1,-1):
                    tempNumber = currentNumber + numList1[n]
                    print("tempNumber 1 = ",tempNumber)
                    if tempNumber <= target:
                        numList1.append(tempNumber)
                        if numList1 in self.record:
                            break
                        length1 = self.findShortestChainHelper(target, numList1, minLength)
                        print("length1 2 = ",length1)
                        break
                    else:
                        continue

        if numListLen >= 2:
            tempNumber = currentNumber + numList2[numListLen-2]
            if tempNumber <= target:
                numList2.append(tempNumber)
                if numList2 not in self.record:
                
                    length2 = self.findShortestChainHelper(target, numList2,minLength)
                    print("length2 1 = ",length2)
            else:
                if numListLen >= 3:
                    for n in range(numListLen-2,-1,-1):
                        tempNumber = currentNumber + numList2[n]
                        print("tempNumber 2 = ",tempNumber)
                        if tempNumber <= target:
                            numList2.append(tempNumber)
                            if numList2 in self.record:
                                break
                            length2 = self.findShortestChainHelper(target, numList2,minLength)
                            print("length2 2 = ",length2)
                            break
                        else:
                            continue
       
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
    additionChain.findShortestChain(133)
