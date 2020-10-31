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

/**
 * add a new thread to the board
 *
 * @return {Promise<boolean>} success post or not
 */
const postThread = async () => {
    try {
        const shortname = window.location.pathname.slice(1)
        const callboard = await fetch(`${api}/boards?shortname=${shortname}`)
        const databoard = await callboard.json()
        const currentBoard = databoard[0].id

        const newThread = document.getElementById('new-thread').value
        const thread = {
            'description': newThread ,
            'board': currentBoard
        }

        const call = await fetch(`${api}/threads`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json;charset=utf-8'
            },
            body: JSON.stringify(thread)
        })
        const data = await call.json()
        if (call.status === 200) {
            return data
        }
    } catch (error) {
        console.error(error)
    }
}

/**
 * call to api to get notifications
 * @return {Promise<Array>}
 */
const getNotifs = async () => {
    try {
        const call = await fetch(`${api}/notifications`)
        const data = await call.json()
        if (call.status === 200) {
            return data
        }
    } catch (error) {
        console.error(error)
    }
}


export { getBoards, getThreads, postThread, getNotifs }
