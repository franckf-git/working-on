import { getBoards, getThreads, getNotifs } from './apiCalls.js'
import { app } from './app.js'

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
        node.title = board.name
        node.id = 'board-link'
        const textnode = document.createTextNode(board.name)
        node.appendChild(textnode)
        subtitle.appendChild(node)
        section.appendChild(subtitle)
    })
}

/**
 * get notif and display it
 * @return {} HTML element
 */
const addNotif = async () => {
    const notifs = await getNotifs()

    const notifcard = document.createElement('div')
    notifcard.classList.add('card')
    notifcard.classList.add('notifications')

    const headercard = document.createElement('header')
    headercard.classList.add('card-header')

    const titlecard = document.createElement('p')
    titlecard.classList.add('card-header-title')
    const messagenotif = document.createTextNode(notifs[0].message)
    titlecard.appendChild(messagenotif)

    headercard.appendChild(titlecard)
    notifcard.appendChild(headercard)

    const footercard = document.createElement('footer')
    footercard.classList.add('card-footer')

    const buttoncard = document.createElement('a')
    buttoncard.classList.add('card-footer-item')
    const getit = document.createTextNode('Get it')
    buttoncard.appendChild(getit)

    footercard.appendChild(buttoncard)
    notifcard.appendChild(footercard)

    app.appendChild(notifcard)

    const getitbutton = document.querySelector('.card-footer-item')
    const notificationcard = document.querySelector('.card')
    getitbutton.addEventListener('click', () => { notificationcard.classList.add('is-hidden') })
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

/**
 * get threads from board shortname and put them in links
 * @param {String} shortname
 * @return {} HTML element
 */
const addThreadsLinks = async (shortname) => {
    const threads = await getThreads(shortname)
    const section = document.createElement('section')
    section.classList.add('section')
    app.appendChild(section)

    const inputThread = document.createElement('input')
    inputThread.classList.add('input')
    inputThread.type = 'text'
    inputThread.id = 'new-thread'
    inputThread.placeholder = 'Add a new thread'
    const inputThreadControl = document.createElement('div')
    inputThreadControl.classList.add('control')
    inputThreadControl.appendChild(inputThread)

    const buttonThread = document.createElement('a')
    buttonThread.classList.add('button')
    buttonThread.id = 'button-new-thread'
    const textButtonThread = document.createTextNode('Add')
    buttonThread.appendChild(textButtonThread)
    const buttonThreadControl = document.createElement('div')
    buttonThreadControl.classList.add('control')
    buttonThreadControl.appendChild(buttonThread)

    const newThreadField = document.createElement('div')
    newThreadField.classList.add('field')
    newThreadField.classList.add('has-addons')
    newThreadField.appendChild(inputThreadControl)
    newThreadField.appendChild(buttonThreadControl)

    section.appendChild(newThreadField)

    threads.forEach((thread) => {
        const subtitle = document.createElement('h2')
        subtitle.classList.add('subtitle')
        const node = document.createElement('a')
        node.href = `${window.location.pathname}/${thread.id}`
        const textnode = document.createTextNode(thread.description)
        node.appendChild(textnode)
        subtitle.appendChild(node)
        section.appendChild(subtitle)
    })

}

export { addBoardsLinks, addThreadsLinks, header, addNotif }
