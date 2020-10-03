import { getBoards } from './apiCalls.js'

const app = document.getElementById('app')

/**
 * get boards and put them in links
 * @return {} HTML element
 */
const addBoardsLinks = async () => {
    const boards = await getBoards()
    const section = document.createElement('section')
    section.classList.add('section')
    app.appendChild(section)

    boards.forEach((board) => {
        const subtitle = document.createElement('h2')
        subtitle.classList.add('subtitle')
        const node = document.createElement('a')
        node.href = board.shortname
        const textnode = document.createTextNode(board.name)
        node.appendChild(textnode)
        subtitle.appendChild(node)
        section.appendChild(subtitle)
    })

}

/**
 * create header
 * @return {} HTML element
 */
const header = () => {
    const title = 'Texts Boards'
    const section = document.createElement('section')
    section.classList.add('hero')
    section.classList.add('is-dark')

    const heroBody = document.createElement('div')
    heroBody.classList.add('hero-body')

    const heroContainer = document.createElement('div')
    heroContainer.classList.add('container')

    const heroTitle = document.createElement('h1')
    heroTitle.classList.add('title')
    const heroTitleText = document.createTextNode(title)

    heroTitle.appendChild(heroTitleText)
    heroContainer.appendChild(heroTitle)
    heroBody.appendChild(heroContainer)
    section.appendChild(heroBody)

    app.appendChild(section)

}

export { addBoardsLinks, header }
