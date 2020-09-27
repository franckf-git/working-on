import { getBoards } from './apiCalls.js'

getBoards().then((value) => {
    console.log(value)
})
