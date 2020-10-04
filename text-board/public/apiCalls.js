const api = 'http://127.0.0.1:5500'

/**
 * call to api to get all boards
 * @return {Promise<Array>} list of the boards
 */
const getBoards = async () => {
    try {
        const call = await fetch(`${api}/boards`)
        const data = await call.json()
        if (call.status === 200) {
            return data
        }
    } catch (error) {
        console.error(error)
    }
}

/**
 * call to api to get all threads from a board
 * @param {String} shortname
 * @return {Promise<Array>} list of the threads
 */
const getThreads = async (shortname) => {
    try {
        const call = await fetch(`${api}/boards?shortname=${shortname}`)
        const data = await call.json()
        if (call.status === 200) {
            return data[0].threads
        }
    } catch (error) {
        console.error(error)
    }
}

export { getBoards, getThreads }
