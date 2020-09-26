const api = 'http://localhost:1337'

/**
 * call to api to get all boards
 * @return {Array} lsit of the boards
 */

const getBoards = async () => {
    const call = fetch(`${api}/boards`)
    console.log(call);

    const result = await call
    return result
}
getBoards()

export { getBoards }