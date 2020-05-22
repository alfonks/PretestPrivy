package main

import(
	"fmt"
)


func fungsiA(slice []string) []string {
	fungsiMap := make(map[string]struct{}) //membuat hash map
	for _, v := range slice{ //melakukan looping sepanjang array slice
		fmt.Println(v)
	    fungsiMap[v] = struct{}{} //melakukan mapping kedalam variable hashmap dari variable slice
	}
	

	fungsiSlice := make([]string, 0, len(fungsiMap)) //membuat array string dengan length sepanjang map yang telah dibuat
	for v := range fungsiMap { //melakukan looping sepanjang hash map fungsiMap
	    fungsiSlice = append(fungsiSlice, v)  //Meng-append nilai dari hashmap kedalam array
	} 
	return fungsiSlice //mereturn nilai dari array fungsiSlice
}

func main(){
	strArray1 := []string{"Japan", "Australia", "Germany"}
	test := fungsiA(strArray1)
	fmt.Println(test)
}
