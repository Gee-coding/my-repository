package main

import (
	"fmt"
	"strings"
)

func main() {

	content := "<dddddddddddddddddddddddddddddd>Hello World!"
	
	fmt.Println(RemoveTagHTML(content,20))

}


func RemoveTagHTML(str string,labelLimit int)(result string){
// remove any string like "<.....>"
	for strings.Contains(str, "<"){
	a := strings.SplitAfterN(str, "<", 2)
	b := strings.SplitAfterN(a[len(a)-1], ">", 2)

	if len(b[0]) > labelLimit {
		return "fund tag label too long. pls make sure "
	}
	if len(b) == 1 {	
		return str
	}
	temp := b[0][0 : len(b[0])-1]
	tagRemove := strings.ReplaceAll(str, "<"+temp+">", "")
	str = tagRemove
}
result = str

return 
}