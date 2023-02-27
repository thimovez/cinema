
module.exports = class FilmDto {
  name;
  genre;
  year;
  constructor(model) {
    this.name = model.name;
    this.genre = model.genre;
    this.year = model.year;
  }
};
