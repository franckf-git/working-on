import express from 'express'
const router = express.Router()

/**
 * users home (public)
 * @route {GET} /
 * @param {object} request
 * @param {object} response
 * @param next
 */
router.get('/', (request: express.Request, response: express.Response, next) => {
  response.render('index', { title: 'Express' })
})

export default router
