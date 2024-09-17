package main

import "fmt"

var inputOpen = map[string] string {
	"(": ")",
		"{": "}",
		"[": "]",
}
var inputClose = map[string] string{
		")": "(",
		"}": "{",
		"]": "[",
}

func main(){
	a1:=[]string{"(", ")", "[", "]"}
	a2:=[]string{"(", "]", "(", "]"}
	a3:=[]string{"(", "(", "(", "(", ")"}
	fmt.Println("A1 = ", perform(a1))
	fmt.Println("A2 = ", perform(a2))
	fmt.Println("A3 = ", perform(a3))
}

func perform(in [] string) bool {
	stack:= []string{}
	for i:=0;i<len(in); i++{
		s:=in[i]
		if _,ok:=inputOpen[s]; ok {
			stack = append(stack, s)
		}

		size:=len(stack)
		if val, ok :=inputClose[s]; ok && stack[size -1] == val   {
			stack = deleteByIndex(stack, size-1)
		}
	}
	return len(stack) == 0
}


func deleteByIndex(in []string, index int) [] string{
	return append(in[:index], in[index+1:]...)
}