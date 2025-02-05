package main

func EachCons(arr []int, n int) [][]int {
    var result [][]int

    for i := 0; i <= len(arr)-n; i++ {

        subset := arr[i : i+n]

        result = append(result, subset)
    }
    
    return result
}
