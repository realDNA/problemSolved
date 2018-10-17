/*
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Example 1:

Input: num1 = "2", num2 = "3"
Output: "6"
Example 2:

Input: num1 = "123", num2 = "456"
Output: "56088"
Note:

The length of both num1 and num2 is < 110.
Both num1 and num2 contain only digits 0-9.
Both num1 and num2 do not contain any leading zero, except the number 0 itself.
You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/

package main

import(
"fmt"
"strconv"
)

func exponent (a,n uint64) uint64  {	
    result := uint64(1)	
    for i := n ; i > 0; i >>= 1 {		
        if i&1 != 0 {			
             result *= a		
        }		
        a *= a	
    }	
    return result
}

func multiply(num1 string, num2 string) string {
    num1_len := len(num1)
    num2_len := len(num2)
    p_length := num1_len + num2_len
    p := make([]uint64, p_length)
    for n2 := num2_len-1; n2 >= 0; n2-- {
        for n1 := num1_len-1; n1 >= 0; n1-- {
            num_mp := (num2[n2]-'0')*(num1[n1]-'0')
            p1 := n1 + n2
            p2 := n1 + n2 + 1
            sum := uint64(num_mp) + p[p2]
            p[p2] = sum % 10
            p[p1] = p[p1] + sum /10
        }
    }
    var multiplier uint64
    multiplier = 1
    var total uint64
    total =0
    total_s := make([]string,0)
    limit := exponent(2,31)
    for i := len(p)-1; i >=0; i-- {
        total += p[i]*multiplier
        if total > limit {
            total_s = append(total_s, strconv.FormatUint(total,10))
            multiplier = 1
            total = 0
        } else { 
            multiplier *= 10
        }
    }
    if total != 0 {
        total_s = append(total_s, strconv.FormatUint(total,10))
    }

    result_s := ""
    for i := len(total_s)-1; i>=0; i-- {
        result_s = result_s + total_s[i]
    }
    if result_s == "" {
        return "0"
    }
    return result_s
}

func main(){
    num1 := "0"
    num2 := "0"
    result := multiply(num1, num2)
    fmt.Println("result = ",result)
}
