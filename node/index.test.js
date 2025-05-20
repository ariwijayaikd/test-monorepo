const request = require('supertest');
const app = require('./app');

describe('GET /', () => {
  it('responds with expected text', async () => {
    const res = await request(app).get('/');
    expect(res.statusCode).toBe(200);
    expect(res.text).toBe('Hello, World! This is a test response from the Node.js server.');
  });
});