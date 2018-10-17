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

func Min(v []int) int {
   m := v[0]
   for _, e := range v[1:] {
      if m==0 && e !=0 {
          m = e
      } else if e < m && e!=0 {
          m = e
      }
   }
   //fmt.Println("min result is ",result)
   if m == 0 {
       return -1
   }
   return m
}

func coinChange(coins []int, amount int) int {
   //fmt.Println("coins = ",coins)
   //fmt.Println("amount = ", amount) 
   
   if amount ==0 {
       return 0
   }

   const Maxint = -1
   dp := make([]int, amount+1)
   dp[0] = 0
   for i:=1; i<=amount; i++ {
       //fmt.Println(i)
       dp[i] = Maxint
   }

   //for i:=1; i<=amount; i++ {
   //    fmt.Println("dp i = ",dp[i])
   //}

   for i:=1; i<= amount; i++ {
       var store []int
       for _,v := range coins {
           if i-v >= 0 {
               store = append(store, dp[i-v]+1)
               //fmt.Println("store = ",store)
           }
       } 
       //fmt.Println("store = ", store)
       if len(store) != 0 {
           dp[i] = Min(store)
       }
       if dp[i] == 0 {
           dp[i] = Maxint
       }
       //fmt.Print(i, dp[i])
       //fmt.Println()
   }
   
   if dp[amount] == Maxint || dp[amount] == 0 {
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

