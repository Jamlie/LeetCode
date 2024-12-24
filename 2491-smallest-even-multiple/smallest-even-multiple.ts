function smallestEvenMultiple(n: number): number {
    return isEven(n) ? n : n * 2;
};

function isEven(n: number): boolean {
    return n % 2 === 0;
}