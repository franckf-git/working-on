import { getBoards } from './apiCalls.js'

const app = document.getElementById('app')

const addBoardsLinks = async () => {
const boards = await getBoards()
    const section = document.createElement('section')
    section.classList.add('section')
    app.appendChild(section)

boards.forEach((board)=>{
    const subtitle = document.createElement('h2')
    subtitle.classList.add('subtitle')
    const node = document.createElement('a')
    node.href=board.shortname
    const textnode = document.createTextNode(board.name)
    node.appendChild(textnode)
    subtitle.appendChild(node)
    section.appendChild(subtitle)
})

}

export { addBoardsLinks }
