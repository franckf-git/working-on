import { getBoards } from "./app";

test('Exist', () => {
    expect(getBoards()).toBeDefined()
})
//@types/jest