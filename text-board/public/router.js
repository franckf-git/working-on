import { app } from './app.js'
import { getThreads } from './apiCalls.js'

/**
 * click boards - change url and clear page
 * @return {} HTML element
 */
const clickBoard = () => {
    app.addEventListener('click', async (event) => {
        if (event.target.id == 'board-link') {
            event.preventDefault()
            history.pushState({}, '', event.target.href)
            app.innerHTML = ''
            const shortname = window.location.pathname.slice(1)
            const threads = await getThreads(shortname)
            console.log(threads)
        }
    })
}

export { clickBoard }
