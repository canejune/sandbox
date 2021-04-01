var testStatus = {
  'name':'',
  'request':''
}

function testStatusRouter(req, res) {
  res.json(testStatus);
}

module.exports = {
  testStatus, testStatusRouter
}