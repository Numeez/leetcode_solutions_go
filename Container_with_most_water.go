package leetcodesolutionsgo
func maxArea(height []int) int {
    result:=0
    left:=0
    right:=len(height)-1
    for left<right{
        area:=0
        if height[right]>height[left]{
            area = height[left]*(right-left)
            left++
        }else{
            area = height[right]*(right-left)
            right--
        }
        if area>result{
            result=area
        }
    }
    return result

}

func findMin(a ,b int) int {
    if a<b{
        return a
    }else{
        return b
    }
}
func findMax(a,b int) int{
    if a>b{
        return a
    }else{
        return b
    }
}