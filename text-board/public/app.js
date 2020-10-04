console.log(`
███████╗██████╗  █████╗ ███╗   ██╗ ██████╗██╗  ██╗███████╗
██╔════╝██╔══██╗██╔══██╗████╗  ██║██╔════╝██║ ██╔╝██╔════╝
█████╗  ██████╔╝███████║██╔██╗ ██║██║     █████╔╝ █████╗  
██╔══╝  ██╔══██╗██╔══██║██║╚██╗██║██║     ██╔═██╗ ██╔══╝  
██║     ██║  ██║██║  ██║██║ ╚████║╚██████╗██║  ██╗██║     
╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝╚═╝  ╚═╝╚═╝     
`)
import { addBoardsLinks, header } from './createDOM.js'
import { clickBoard } from './router.js'
export const app = document.getElementById('app')

if (window.location.pathname === '/') {
    header()
    addBoardsLinks()
    clickBoard()
}
