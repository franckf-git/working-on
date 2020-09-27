//const { getBoards } = require('./../app')
import { getBoards } from './../app'

test('Exist', () => {
    expect(getBoards()).toBeDefined()
})

test('Get at least one board', () => {
    const firstBoard = getBoards()[0]
    expect(firstBoard.id).toBe('1')
})
