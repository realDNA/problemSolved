/*
Ordered fractions
Problem 71 
Consider the fraction, n/d, where n and d are positive integers. If n<d and HCF(n,d)=1, it is called a reduced proper fraction.

If we list the set of reduced proper fractions for d ≤ 8 in ascending order of size, we get:

1/8, 1/7, 1/6, 1/5, 1/4, 2/7, 1/3, 3/8, 2/5, 3/7, 1/2, 4/7, 3/5, 5/8, 2/3, 5/7, 3/4, 4/5, 5/6, 6/7, 7/8

It can be seen that 2/5 is the fraction immediately to the left of 3/7.

By listing the set of reduced proper fractions for d ≤ 1,000,000 in ascending order of size, find the numerator of the fraction immediately to the left of 3/7.
*/

/*
b/a < 3/7
1. 3a - 7b =1 -> b = (3a-1) / 7
2. a,b <= 1e6
3. HCF(a,b) = 1
4. a,b are max number each in 1e6  
*/

package main

import(
"fmt"
)
func isInt(a float64) bool {
    if a == float64(int64(a)) {
        return true
    }
    return false
}
 

func orderedFractions(n int) [2]int64 {
   var b float64
   var aMax int64
   var bMax int64
   var final [2]int64

   for a := 0.0; a<=1e6; a++ {
       b = (3*a - 1) / 7
       
       aInt64 := int64(a)
       bInt64 := int64(b) 
       if isInt(b) && (aInt64 % bInt64 != 0){
           aMax = aInt64
           bMax = bInt64
       }
   }
   final[0] = aMax
   final[1] = bMax
   
   return final   
}

func main() {
    n := 1000000
    final := orderedFractions(n)
    fmt.Println("a = ", final[0])
    fmt.Println("b = ", final[1])
}
