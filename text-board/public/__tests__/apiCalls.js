import { enableFetchMocks } from 'jest-fetch-mock'
enableFetchMocks()
import { getBoards } from "../apiCalls"

test('Exist', () => {
    expect(getBoards()).toBeDefined()
})

describe('Testing /boards API', () => {
    beforeEach(() => {
        fetch.resetMocks()
    })
    test('Get at least one board', async () => {
        await fetch.mockResponse('[{"id":1,"name":"technology"},{"id":2,"name":"fetchmock"}]')
        const firstBoard = await getBoards()
        expect(firstBoard[0].id).toBeGreaterThanOrEqual(1)
        expect(firstBoard[0].name).toMatch(/[a-zA-Z]/)
    })
})
