import { app } from './app.js'
import { addThreadsLinks } from './createDOM.js'

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
            addThreadsLinks(shortname)
        }
    })
}

export { clickBoard }
