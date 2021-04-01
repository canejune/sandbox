const express = require('express');
const morgan = require('morgan');
const app = express();
const port = 3000;
const {config} = require('./config.js');
const {infoRouter} = require('./info.js');
const {testRequestRouter} = require('./testRequest.js');
const {testStatusRouter} = require('./testStatus.js');

app.use(morgan('dev'));
app.use(express.json()) // for parsing application/json
app.use(express.urlencoded({ extended: true })) // for parsing application/x-www-form-urlencoded

app.get('/', (req, res) => {
  res.json(config);
});

app.get('/api/info', infoRouter);
app.post('/api/test-request', testRequestRouter);
app.get('/api/test-status', testStatusRouter);

app.listen(port, () => {
  console.log(`app listening at http://localhost:${port}`);
});