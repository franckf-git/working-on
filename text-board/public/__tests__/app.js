//const { getBoards } = require('./../app')
import { getBoards } from './../app'

test('Exist', () => {
    expect(getBoards()).toBeDefined()
})

test('Get at least one board', async () => {
    const firstBoard = await getBoards()
    expect(firstBoard[0].id).toBeGreaterThanOrEqual(1)
})
