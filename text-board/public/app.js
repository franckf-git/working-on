console.log(`
███████╗██████╗  █████╗ ███╗   ██╗ ██████╗██╗  ██╗███████╗
██╔════╝██╔══██╗██╔══██╗████╗  ██║██╔════╝██║ ██╔╝██╔════╝
█████╗  ██████╔╝███████║██╔██╗ ██║██║     █████╔╝ █████╗  
██╔══╝  ██╔══██╗██╔══██║██║╚██╗██║██║     ██╔═██╗ ██╔══╝  
██║     ██║  ██║██║  ██║██║ ╚████║╚██████╗██║  ██╗██║     
╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝╚═╝  ╚═╝╚═╝     
`)
import { addBoardsLinks, header } from './createDOM.js'
import { clickBoard, addNewThread } from './router.js'
import { getNotifs } from './apiCalls.js'
export const app = document.getElementById('app')

if (window.location.pathname === '/') {
    header()
    addBoardsLinks()
    clickBoard()
    addNewThread()
}

const testnotif = async () => {
    const api = await getNotifs()
    console.log(api)
}

testnotif()
