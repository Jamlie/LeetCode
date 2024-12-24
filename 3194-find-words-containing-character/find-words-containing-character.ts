function findWordsContaining(words: string[], x: string): number[] {
    return words
        .map((word, i) => (word.includes(x) ? i : NaN))
        .filter((v) => !isNaN(v));
}
