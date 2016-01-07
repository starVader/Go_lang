package main 

import "fmt"


func Filter(data []int,f func(int) bool, index int, result []interface{}) []interface{}{
	if index > (len(data)-1){
		fmt.Println("Done with recursion")
		return result
	}
	value:= f(data[index])
	if value{
		result = append(result,data[index])
	}
	index++
	return Filter(data[:],f,index,result)
}

func main(){
	even:= func (x int) bool{
		if x%2 == 0 {
			return true
		}
		return false
	}
	arr:= []int{}
	fmt.Println(arr)
	result:= make([]interface{},0)
	fmt.Println(Filter(arr[:],even,0,result))

}