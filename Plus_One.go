/*
You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer.
 The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.


Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].


Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].

digits = [9,9]
Output: [1,0,0]

*/

package main

import (
    "fmt"
)


func plusOne(digits []int) []int {
    res := make([]int, 0)

    if len(digits) == 0 {
        fmt.Println("Slice is empty")
    }

    if len(digits) == 1 {
        if digits[len(digits) - 1]  == 9 {
            num_f := 0
            num_s := 1
            res = append(res, num_s,num_f)
            return res
        } else {
            digits[0] += 1 
            res = digits
            return res
        }
    } else {
        if digits[len(digits) - 1]  == 9 {
            res = []int{0}
            return append(plusOne(digits[:len(digits) - 1]), res...)
        } else {
            digits[len(digits) - 1] = digits[len(digits) - 1] + 1
            res = digits
        }
    }
    
    
    return res

}


func main() {
    res := []int{1,2,3}
    fmt.Println(plusOne(res))

}
