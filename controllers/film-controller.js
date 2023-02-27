const sortByYear = require('../service/film-service');

/** FilmController controls the behavior of films */
class FilmController {
  async filterByGenre(_req, res) {
    try {
      const films = await sortByYear();
      console.log(films);
      return res.json(films);
    } catch (e) {
      console.log(e);
    }
  }
}

module.exports = new FilmController();
