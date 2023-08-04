package main

import "fmt"

//function to convert characters to upper case
func ToUpper(s string) string {
	b := []byte(s)
	for i, c := range b {
		if c >= 'a' && c <= 'z' {
			b[i] = c - ('a' - 'A')
		}
	}
	return string(b)
}

func main() {

	var array = [...]string{"sachin", "shewaga", "dhoni", "Virat"}
	var length = len(array)

	for i := 0; i < length; i++ {
		rank := i + 1
		fmt.Println(rank, "Player", ToUpper(array[i]))
	}

}
