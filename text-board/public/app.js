const api = 'http://localhost:1337'

/**
 * call to api to get all boards
 * @return {Array} lsit of the boards
 */
const getBoards = async () => {
    const call = fetch(`${api}/boards`)

    const result = await call
    console.log(result)
    return result
}

module.exports = { getBoards }
// export { getBoards }