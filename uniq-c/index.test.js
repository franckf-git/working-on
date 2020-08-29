const { uniq } = require('./index')

it('should return an object', () => {
    const result = uniq('a string with a lot of words')
    expect(typeof result).toBe('object')
})
