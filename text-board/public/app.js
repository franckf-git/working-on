console.log(`
███████╗██████╗  █████╗ ███╗   ██╗ ██████╗██╗  ██╗███████╗
██╔════╝██╔══██╗██╔══██╗████╗  ██║██╔════╝██║ ██╔╝██╔════╝
█████╗  ██████╔╝███████║██╔██╗ ██║██║     █████╔╝ █████╗  
██╔══╝  ██╔══██╗██╔══██║██║╚██╗██║██║     ██╔═██╗ ██╔══╝  
██║     ██║  ██║██║  ██║██║ ╚████║╚██████╗██║  ██╗██║     
╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝╚═╝  ╚═╝╚═╝     
`)
import { addBoardsLinks, header } from './createDOM.js'


if (window.location.pathname === '/') {
    header()
    addBoardsLinks()
}

// click link
document.querySelector('#app').addEventListener('click', (event) => {
    if (event.target.id == 'board-link') {
        event.preventDefault()
        history.pushState({}, '', event.target.href)
    }
})
