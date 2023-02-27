const fs = require("fs");
const csv = require("csv-parser");
const { resolve } = require("path");
const { rejects } = require("assert");

function quickSort(arr) {
  if (arr.length < 2) return arr;
  let pivot = arr[0].year;
  const left = [];
  const right = [];
    
  for (let i = 1; i < arr.length; i++) {
    if (pivot > arr[i].year) {
      left.push(arr[i]);
    } else {
      right.push(arr[i]);
    }
  }

  return quickSort(left).concat(arr[0], quickSort(right));
}

const objectToCsv = function (data) {

  const csvRows = [];

  /* Get headers as every csv data format
  has header (head means column name)
  so objects key is nothing but column name
  for csv data using Object.key() function.
  We fetch key of object as column name for
  csv */
  const headers = Object.keys(data[0]);

  /* Using push() method we push fetched
     data into csvRows[] array */
  csvRows.push(headers.join(','));

  // Loop to get value of each objects key
  for (const row of data) {
      const values = headers.map(header => {
          const val = row[header]
          return `"${val}"`;
      });

      // To add, separator between each value
      csvRows.push(values.join(','));
  }

  /* To add new line for each objects values
     and this return statement array csvRows
     to this function.*/
  return csvRows.join('\n');
};

const sortByYear = async () => new Promise((resolve, reject) => {
    let movies = [];
    fs.createReadStream('./movies.csv')
      .pipe(csv())
      .on("data", (data) => {
        const film = {
          film: data.film,
          genre: data.genre,
          year: data.year
        };

        movies.push(film);
    })  
    .on("end", () => {
        console.log("done!");
        let sortedMovies = quickSort(movies);
        let csv = objectToCsv(sortedMovies);
        fs.writeFileSync('./some.csv', csv);
        resolve(csv);
    })
    .on('error', function (error) {
        reject(error)
    });
});

module.exports = sortByYear;