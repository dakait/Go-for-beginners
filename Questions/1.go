package main

import (
	"encoding/json"
	"fmt"
)

//Assume these structs were used by someone to marshal the data
type Test struct {
	Id int64 `json:"id"`
	Msg string `json:"msg"`
}

type Test2 struct {
	Yet string `json:"yet"`
	Pkg string `json:"pkg"`
}



func main() {
	//this part marshals the data
	t1 := Test{Id:1, Msg:"No message haha"}

	t2 := Test2{Yet:"Yet another yum", Pkg:"yum-12.34"}
	
	newj, _ := json.Marshal(map[string]interface{}{"t1":t1, "t2":t2})
	
	//you received this data(newj) now you'll unmarshal it without using struct
	
	var tt map[string]interface{}
	//unmarshalling data
	json.Unmarshal(newj, &tt)
	// type chaecking and returning relevent type value
	m , ok := tt["t1"].(map[string]interface{})

	if ok {
		for i, _ := range m {
			fmt.Println(i, ":" ,m[i])
		}
	}
}
