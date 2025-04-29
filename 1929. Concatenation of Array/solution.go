func getConcatenation(nums []int) []int {
    
    // new slice for a secondary array or slice
    answer := make([]int, 0, len(nums) * 2)

    // appending the first nums slice
    answer = append(answer, nums...) 

    // appending again the first nums
    answer = append(answer, nums...) 
     
    // now this is concatenated slice and this is the solution

    return answer
}
