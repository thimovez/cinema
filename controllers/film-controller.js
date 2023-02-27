const sortByYear = require('../service/film-service');

class FilmController {

    async filterByGenre(req, res) {
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