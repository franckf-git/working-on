const { processNumber, collatzProcess } = require('./index')

it('should be divide by 2', () => {
    expect(processNumber(8)).toBe(4)
})

it('should be multiply by 3, and plus 1', () => {
    expect(processNumber(3)).toBe(10)
})

it('should return an array', () => {
    expect(collatzProcess(8)).toContain(8)
})

it('should return a full of collatz array', () => {
    const expectArray = [8,4,2,1]
    expect(collatzProcess(8)).toStrictEqual(expectArray)
})
