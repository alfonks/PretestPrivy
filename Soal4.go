package main

import(
	"fmt"
)


func fungsiA(slice []string) []string {
	fungsiMap := make(map[string]struct{})
	for _, v := range slice{ 
	    fungsiMap[v] = struct{}{} 
	} 
	 
	fungsiSlice := make([]string, 0, len(fungsiMap)) 
	for v := range fungsiMap { 
	    fungsiSlice = append(fungsiSlice, v) 
	}
	fmt.Println(fungsiSlice) 
	return fungsiSlice 
}

func main(){
	strArray1 := []string{"Japan", "Australia", "Germany"}
	fungsiA(strArray1)
}
