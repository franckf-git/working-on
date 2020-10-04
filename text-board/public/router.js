import { app } from './app.js'

/**
 * click boards - change url and clear page
 * @return {} HTML element
 */
const clickBoard = () => {
    app.addEventListener('click', (event) => {
        if (event.target.id == 'board-link') {
            event.preventDefault()
            history.pushState({}, '', event.target.href)
            app.innerHTML = ''
            console.log(window.location.pathname)
        }
    })
}

export { clickBoard }
