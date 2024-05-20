package leetcodesolutionsgo
func isAcronym(words []string, s string) bool {
    if len(words)!=len(s){
        return false
    }
    result:=""
    for _,word:= range words{
        result+=string(word[0])
    }
    if result==s{
        return true
    }
    return false

    
}