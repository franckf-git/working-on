const fetch = require('node-fetch')
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

export { getBoards }
