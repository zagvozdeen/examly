const request = require('supertest')
const { fakerRU: faker } = require('@faker-js/faker')

const whatsInside = done => {
    return (err, res) => {
        if (err) {
            done(err)
            return
        }

        console.log(res.body)

        done()
    }
}

describe('/auth/register', function () {
    this.timeout(5000)

    it('method should not allowed', function (done) {
        request("http://127.0.0.1:8080")
            .get("/api/v2/auth/register")
            .expect(405, done)
    })

    it('method should bad request with empty body', function (done) {
        request("http://127.0.0.1:8080")
            .post("/api/v2/auth/register")
            .expect(400)
            .expect('Content-Type', 'application/json')
            .expect(/{"message":".*"}/, done)
    })

    const email = faker.internet.email()

    it('method should success created', function (done) {
        request("http://127.0.0.1:8080")
            .post("/api/v2/auth/register")
            .send({
                first_name: faker.person.firstName(),
                last_name: faker.person.lastName(),
                email: email,
                password: 'password',
                password_confirmation: 'password',
            })
            .expect(201, done)
    })

    it('method should bad request', function (done) {
        request("http://127.0.0.1:8080")
            .post("/api/v2/auth/register")
            .send({
                first_name: faker.person.firstName(),
                last_name: faker.person.lastName(),
                email: email,
                password: 'password',
                password_confirmation: 'password',
            })
            .expect(500)
            .expect('Content-Type', 'application/json')
            .expect(/{"message":".*"}/)
            .end(whatsInside(done))
    })
})