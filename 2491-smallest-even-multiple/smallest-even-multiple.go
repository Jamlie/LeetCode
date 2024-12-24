func smallestEvenMultiple(n int) int {
    if isEven(n) {
        return n
    }

    return n * 2
}

func isEven(n int) bool {
    return n % 2 == 0;
}