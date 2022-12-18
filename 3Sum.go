/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k,
 and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.


Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation: 
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.


Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.


Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
*/


package main

import (
    "fmt"
)

func partition(arr []int, low, high int) ([]int, int) {
    pivot := arr[high]
    i := low
    for j := low; j < high; j++ {
        if arr[j] < pivot {
            arr[i], arr[j] = arr[j], arr[i]
            i++
        }
    }
    arr[i], arr[high] = arr[high], arr[i]
    return arr, i
}


func quickSort(arr []int, low, high int) []int {
    if low < high {
        var p int
        arr, p = partition(arr, low, high)
        arr = quickSort(arr, low, p-1)
        arr = quickSort(arr, p+1, high)
    }
    return arr
}


func quickSortStart(arr []int) []int {
    return quickSort(arr, 0, len(arr)-1)
}


func threeSum(nums []int) [][]int {
    var res [][]int
    nums = quickSortStart(nums)

    for i := 0; i < len(nums) - 2; i++ {
        
        if i != 0 && nums[i] == nums[i - 1]{
            continue
        }
        
        for j,k := i + 1,len(nums) - 1; j < k;{
            
            if j != i + 1 && nums[j] == nums[j - 1] {
                j++
                continue
            }
            
            if k < len(nums) - 1 && nums[k] == nums[k + 1]{
                k--
                continue
            }
       
            switch sum := nums[i] + nums[j] + nums[k];{
                case sum == 0:
                    v := []int{nums[i],nums[j],nums[k]}
                    res = append(res,v)
                    j++ 
                    k--
                case sum < 0:
                    j++
                default:
                    k--;
            }
        }
    }
    
    return res
}


func main() {
    nums1 := []int{0,1,1} 
    fmt.Println(threeSum(nums1))
    /*nums2 := []int{0,0,0,0}
    nums3 := []int{0,1,1}*/
   /* fmt.Println(threeSum(nums1))
    fmt.Println(threeSum(nums2))
    fmt.Println(threeSum(nums3))*/
   
}