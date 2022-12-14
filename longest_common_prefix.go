/*
14. Longest Common Prefix
easy


Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

*/


package main

import (
	"fmt"
	"strings"
)


func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    var prefix = strs[0]
    for _, word := range strs {
        for !strings.HasPrefix(word, prefix){
            prefix = prefix[0:len(prefix) - 1]
        }
    }
    if len(prefix) == 0 {
        return ""
    }
    return prefix
}


func main() {
	strs := []string{"slovo", "sloboda", "sleep"}
	fmt.Println(longestCommonPrefix(strs))
}
