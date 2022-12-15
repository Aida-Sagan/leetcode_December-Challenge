package main

import "fmt"


type Stack []rune


func (s *Stack) isEmpty() bool {
    return len(*s) == 0
}

func isValid(s string) bool {
    var stack Stack

    for _, ch := range s {

        if ch == '(' || ch == '{' || ch == '[' {
            stack = append(stack, ch)

        } else if !stack.isEmpty() {
            top := stack[len(stack) - 1]
            
            if top == '(' && ch == ')' || top == '{' && ch == '}' || top == '[' && ch == ']' {
                stack = stack[:len(stack) - 1]
            } else {
                return false
            }
        } else {
            return false
        }
    }
    return stack.isEmpty()
}


func main() {
   s := "()[]{}"

   fmt.Println(isValid(s))
} 
