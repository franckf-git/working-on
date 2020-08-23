const { addTwo, fibo, binet } = require('./index')

it('should sum 2 numbers ', () => {
    expect(addTwo(2,3)).toBe(5)
})

it('should return an array ', () => {
    expect(fibo(2)).toStrictEqual([1,1])
})

it('should process as long as need ', () => {
    const resultion = [1,1,2,3,5,8,13,21,34]
    expect(fibo(9)).toStrictEqual(resultion)
})

