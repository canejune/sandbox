var {testStatus} = require('./testStatus.js');

function testRequestRouter(req, res) {
  console.log(req.body);
  testStatus.name = req.body.name;
  testStatus.request = req.body.request;
  res.status(201).send(req.body);
}

module.exports = { testRequestRouter };