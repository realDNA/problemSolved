'''
The Ackermann function
Problem 282 
For non-negative integers m, n, the Ackermann function A(m,n) is defined as follows:
A(m,n)={n+1A(m−1,1)A(m−1,A(m,n−1)) if m=0 if m>0 and n=0 if m>0 and n>0
For example A(1,0)=2, A(2,2)=7 and A(3,4)=125.

Find ∑n=06A(n,n) and give your answer mod 148.
'''

class AckermannSolution(object):
    def __init__(self):
        self.record = {}  # prevent maximum recursion depth exceeded
        print("self.record = ", self.record)
    
    def ackermannFormula(self, m, n):
        if (m,n) in self.record.keys():
                return self.record[(m,n)]

        if m==0:
            result = n+1
        elif m>0 and n==0:
            result = self.ackermannFormula(m-1, 1)
        else:
            result = self.ackermannFormula(m-1, self.ackermannFormula(m, n-1))

        self.record[(m,n)] = result
        return result
    
    def test(self,m,n):
        for i in range(m,-1,-1):
            for j in range(n,-1,-1):
               
                if j==0:
                    n = 1
                print(f"{i},{j}")
               
                if i==0:
                    self.record[(i,j)]=j+1
                    print(f"to end result = {j+1}")
                    break

if __name__=="__main__":
    ackermannSolution = AckermannSolution()
    sumResult = 0
    ackermannSolution.test(3, 3)
    #sumResult = ackermannSolution.ackermannFormula(3, 29)
    #print("sumResult = ",sumResult)
    print("self.record = ", ackermannSolution.record)
    #for i in range(0,7):
    #    sumResult = ackermannSolution.ackermannFormula(i, i)
    #print("sumResult = ",sumResult)
