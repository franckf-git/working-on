import request from 'supertest'
import app from './../app'

test("basic route users", async () => {
    const response = await request(app).get('/users')
    expect(response.status).toBe(200)
    expect(response.status).not.toBe(404)
    expect(response.type).toBe('text/html')
    expect(response.headers).toBeDefined()
    expect(response.text).toBe('respond with a resource')
})
