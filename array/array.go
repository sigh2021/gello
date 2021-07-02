package myarray

func ArraySum(a []int) (sum int) {
    for _, v := range a {
        sum += v
    }
    return
}

func SumAll(numbersToSum ...[]int)(result []int) {
    //result = make([] int, length)
    for _,numbers := range numbersToSum {
        result = append(result,ArraySum(numbers))
    }
    return
}
