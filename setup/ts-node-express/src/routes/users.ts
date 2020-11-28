import express from 'express'
const router = express.Router()

/**
 * users page (public)
 * @route {GET} /users
 * @param {object} request
 * @param {object} response
 * @param next
 */
router.get('/', (request: express.Request, response: express.Response, next) => {
  response.send('respond with a resource')
})

export default router
