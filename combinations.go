package main

import "fmt"


var results [][]int
var path []int

func backtrack(nums []int, idx, count int) {
    if count == 0 {
        // found the combination
        result := make([]int, len(path))
        copy(result, path)
        results = append(results, result)
        return
    }
    for i := idx; i < len(nums); i++ {
        path = append(path, nums[i])
        backtrack(nums, i+1, count-1)
        path = path[:len(path) - 1]
    }

}

func combine(n int, k int) [][]int {
    results = make([][]int, 0)
    path = make([]int, 0)
    nums := make([]int, n)
    for i := 1; i <= n; i++ {
        nums[i-1] = i
    }
    backtrack(nums, 0, k)
    return results
}

func main() {
    n := 4
    k := 2

    fmt.Println(combine(n, k))
}