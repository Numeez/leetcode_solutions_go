package leetcodesolutionsgo
func reverse(x int) int {
    MIN := -2147483648
    MAX := 2147483647

    res:=0
    for x!=0{
        digit:= x%10
        x=x/10

        if (res>MAX/10) || (res==MAX/10 && digit>=MAX%10){
            return 0
        }
        if (res<MIN/10) || (res==MIN/10 && digit<=MIN%10){
            return 0
        }
        res = (res*10)+digit
    }
    return res
    
}