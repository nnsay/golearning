// setTimeout((str) => {
//   console.log(`${str}  main setTimeout`)
// }, 100, 'hello');

// setImmediate((str) => {
//   console.log(`${str}  main setImmediate`)
// }, 'hello');
const fs = require('fs')

fs.readFile('./README.md', (err, data) => {
  setTimeout((str) => {
    console.log(`${str}  setTimeout`)
  }, 0, 'hello')

  setImmediate((str) => {
    console.log(`${str}  setImmediate`)
  }, 'hello')
});
