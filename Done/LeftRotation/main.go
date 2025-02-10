package main

func rotateLeft(d int32, arr []int32) []int32 {
    n := int32(len(arr))
    if n == 0 || d <= 0 {
        return arr
    }

    d = d % n 
    for i := int32(0); i < d; i++ {
        first := arr[0]
        arr = arr[1:]
        arr = append(arr, first)
    }

    return arr
}