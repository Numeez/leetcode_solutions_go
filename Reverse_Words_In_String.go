package leetcodesolutionsgo

import(
    "strings"
	"fmt"
)
func reverseWords(s string) string {
    result:=""
    words:=strings.Fields(s)
    for i:=len(words)-1;i>=0;i--{
        result = result+words[i]+" "
    }
    fmt.Println(words)
   
    return ( result[:len(result)-1])
   
    
}