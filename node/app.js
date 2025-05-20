const express = require('express');
const app = express();

app.get('/', (req, res) => {
  res.send('Hello, World! This is a test response from the Node.js server.');
});

module.exports = app;