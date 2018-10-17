/*
You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

Example 1:

Input: coins = [1, 2, 5], amount = 11
Output: 3 
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Note:
You may assume that you have an infinite number of each kind of coin.
*/

package main

import(
"fmt"
)

func coinChange(coins []int, amount int) int {
   
   if amount ==0 {
       return 0
   }

   dp := make([]int, amount+1)
   dp[0] = 0
   for i:=1; i<=amount; i++ {
       dp[i] = -1
   }

   for i:=1; i<= amount; i++ {
       store := -1
       flag := 0
       for _,v := range coins {
           if i-v >= 0 {
               dp_temp := dp[i-v]+1
               if dp_temp != 0 {
                  if flag != 1 {
                      store = dp_temp
                      flag = 1
                  } else {
                      if dp_temp < store {
                         store = dp_temp
                      }
                  }
               }
           }
       } 
       dp[i] = store
   }
   
   if dp[amount] == -1 {
      return -1
   }

   return dp[amount]
}

func main(){
    amount := 6249
    coins := []int{186, 419, 83, 408}
    total := coinChange(coins, amount)    
    fmt.Println("final result = ",total)
}

