/*
26. Remove Duplicates from Sorted Array


Given an integer array nums sorted in non-decreasing order, 
remove the duplicates in-place such that each unique element appears only 
once.
 The relative order of the elements should be kept the same.

*/

package main

import "fmt"


func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    var res int = 1
    for i := 1; i < len(nums); i++ {
        if nums[i-1] != nums[i] {
            nums[res] = nums[i]
            res++
        }
    }
    return res
}


func main(){
   nums := []int{1,1,2,3,4,4,4,5}
   fmt.Println(removeDuplicates(nums))
}