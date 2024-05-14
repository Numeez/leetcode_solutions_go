package leetcodesolutionsgo
import(
	"strings"
)

func intToRoman(num int) string {
	symList := [][]interface{}{
		 {"M", 1000},
		 {"CM", 900},
		 {"D", 500},
		 {"CD", 400},
		 {"C", 100},
		 {"XC", 90},
		 {"L", 50},
		 {"XL", 40},
		 {"X", 10},
		 {"IX", 9},
		 {"V", 5},
		 {"IV", 4},
		 {"I", 1},
	 }
 
	 var res strings.Builder
 
	 for _, symVal := range symList {
		 sym := symVal[0].(string)
		 val := symVal[1].(int)
 
		 if num/val > 0 {
			 count := num / val
			 res.WriteString(strings.Repeat(sym, count))
			 num = num % val
		 }
	 }
 
	 return res.String()
 
	 
 }