import request from 'supertest';
import app from './../app';

test("basic route home", async () => {
    const response = await request(app).get('/')
    expect(response.status).toBe(200)
    expect(response.status).not.toBe(404)
    expect(response.type).toBe('text/html')
    expect(response.headers).toBeDefined()
    expect(response.text).toMatch(/Express/)
})
