package main

import "fmt"


func search(nums []int, target int) int {
    left, right := 0, len(nums) - 1

    if len(nums) == 0 {
        return -1
    }

    for left < right {
        mid_ind := (left + (right - left)/2)
        if nums[mid_ind] == target {
            return mid_ind
        }
        //nums[mid:right]
        if nums[mid_ind] < nums[right] {
            if nums[mid_ind] > target {
                right = mid_ind-1
                continue
            }

            if target <= nums[right] {
                left = mid_ind + 1
            } else {
                right = mid_ind - 1
            }
            continue
        }

    // nums[left:mid] 
        if nums[mid_ind] < target {
            left = mid_ind + 1
            continue
        }

        if nums[left] <= target {
            right = mid_ind - 1
        } else {
            left = mid_ind + 1
        }

        
    }
    if left == right && nums[left] == target {
        return left
    }
    return -1
}

func main(){
    var nums []int = []int{}
    
    target := 5
    fmt.Println(search(nums, target))
}   
