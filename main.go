package main

import(
	"fmt"
	"os"
	"sort"
)

func main(){

	var n int
	var input string
	fmt.Println("Ingresa el tamaño del Slice: ")
	fmt.Scanln(&n)

	var s []string

	for i:=0; i < n; i++{
		fmt.Println("valor string: ")
		fmt.Scan(&input)
		s = append(s,input)
	}


	//Descendente
	sort.Sort(sort.Reverse(sort.StringSlice(s)))

	file, err := os.Create("descendente.txt")
		if err != nil {
			fmt.Println(err)
			return
		}

	defer file.Close()
	
	for j:=0; j < n; j++{
		file.WriteString(s[j])
		file.WriteString("\n")
	} 

	//Ascendente
	sort.Sort(sort.StringSlice(s))
	
	file1, err1 := os.Create("ascendente.txt")
	if err1 != nil {
		fmt.Println(err)
		return
	}
	
	defer file1.Close();
	
	for j:=0; j < n; j++{
		file1.WriteString(s[j])
		file1.WriteString("\n")
	} 

}