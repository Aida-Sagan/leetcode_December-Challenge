package main

import "fmt"

func removeElement(nums []int, val int) int {
    
    if len(nums) == 0 {
        return 0
    }
    
    res := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[res] = nums[i]
            res++
        } 
    }
    return res
}

func main() {
    nums := []int{0,1,2,2,3,0,4,2}
    val := 2

    fmt.Println(removeElement(nums, val))
} 
