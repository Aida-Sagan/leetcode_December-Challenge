package main

import "fmt"

func partitionQuickSort(nums []int, start, end int) ([]int, int) {
    var pivot int = nums[end]
    i := start
    if len(nums) == 0 {
        return nums, 0
    } else {
        for j := start; j < end; j++ {
            if nums[j] < pivot {
                nums[i], nums[j] = nums[j], nums[i]
                i++
            }
        }
        nums[i], nums[end] = nums[end], nums[i]
    }
    return nums, i
}

func quick_sort_do(nums []int, start, end int) []int {
    if start < end {
        var pivot int 
        nums, pivot = partitionQuickSort(nums, start, end)
        nums = quick_sort_do(nums, start, pivot - 1)
        nums = quick_sort_do(nums, pivot + 1, end)
    }
    
    return nums 
}

func quick_sort(nums []int) []int {
    return quick_sort_do(nums, 0, len(nums) - 1)
}

func main() {
    nums := []int{-9,0,-15,52,10}
    fmt.Println(quick_sort(nums))
}