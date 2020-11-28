import request from 'supertest'
import app from './../app'

test("if a 404 if displayed on unknown pages", async () => {
    const response = await request(app).get('/arandompagewithoutroute')
    expect(response.status).toBe(404)
    expect(response.type).toBe('text/html')
    expect(response.headers).toBeDefined()
    expect(response.text).toMatch(/Not.Found/)
})
