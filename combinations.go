package main

import "fmt"

/*
Given two integers n and k, return all possible combinations of k numbers 
chosen from the range [1, n].

You may return the answer in any order.


Input: n = 4, k = 2
Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
Explanation: There are 4 choose 2 = 6 total combinations.
Note that combinations are unordered, i.e., [1,2] and [2,1] 
are considered to be the same combination.

*/
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