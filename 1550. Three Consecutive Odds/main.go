func threeConsecutiveOdds(arr []int) bool {
    var count int = 0

    for i:=range arr{
        if arr[i] % 2 != 0{
            count++
            if(count == 3){
                return true
            }
        }else{
            count = 0;
        }
    }
    
    return false
}
