function findWordsContaining(words: string[], x: string): number[] {
    return words
        .map((word, i) => {
            const doesConatin = word.includes(x);

            return doesConatin ? i : NaN;
        })
        .filter((v) => !isNaN(v));
}