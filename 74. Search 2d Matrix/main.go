func searchMatrix(matrix [][]int, target int) bool {

    for _, row := range matrix{
        low := 0
        high := len(row) - 1
        
        for low <= high{

            mid := low + (high - low) / 2

            if row[mid] == target{
                return true
            }else if row[mid] < target{
                low = mid + 1
            }else if row[mid] > target{
                high = mid - 1
            }
        }
    }

    
    return false
}
