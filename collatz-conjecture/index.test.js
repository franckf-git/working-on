const processNumber = require('./index')

it('should be divide by 2', () => {
  expect(processNumber(8)).toBe(4)
})

it('should be multiply by 3, and plus 1', () => {
  expect(processNumber(3)).toBe(10)
})
