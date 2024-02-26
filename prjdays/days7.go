package prjdays

import "fmt"
func Days7() {
    var slice []string = []string{"How", "are", "you", "my", "friend", "today", "?"} 
    for i := len(slice) - 1 ; i >= 0; i--{
        fmt.Println(slice[i])
    }
}

func DivideSlise() {
    array := [5]int{1,2,3,4,5}
    var slice1 []int
	var slice2 []int
    for _, el := range array{
        if el %2 == 0{
			slice1 = append(slice1, el)
        }else{
		slice2 = append(slice2, el)
		}
        fmt.Println("")
    }
    fmt.Println(slice1, slice2)
}