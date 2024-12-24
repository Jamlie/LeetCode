function maximumWealth(accounts: number[][]): number {
    return Math.max(
        ...accounts.map((account) => account.reduce((v, acc) => v + acc, 0)),
    );
}