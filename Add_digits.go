package leetcodesolutionsgo
func addDigits(num int) int {
    if num <=9{
        return num
    }
    sum := 0
    for num>0{
    last_digit:=num%10
    num = num/10
    sum+=last_digit
    }
    num = sum
   



    return addDigits(sum)


}