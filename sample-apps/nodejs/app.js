const express = require('express')

const app = express()
const httpServer = require('http').createServer(app)

app.get('/', (req, res) => {
    res.send('We gotta go, Hello, We gotta go')
})

httpServer.listen(8081, () => console.log("Hello, hello, Blessed Liked Sunday, Flyer than a runway!"))