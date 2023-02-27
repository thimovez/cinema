const express = require('express');
const filmController = require('./controllers/film-controller');
const app = express();
const port = 3000;

app.get('/sort', filmController.filterByGenre);

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
