//Example code shorthand Javascript Coding Techniques

/*1.Ternary Operator
Longhand : 
*/
const x = 20;
let answer;

if (x > 10) {
    answer = "greater than 10";
} else {
    answer =  "less than 10";
}

//Shorthand : 
const answer = x > 10 ? "greater than 10" : "less than 10";
//you can nest your if statement : 
const answer = x > 10 ? "greater than 10" : x < 5 ? "less than 5" : "between 5 and 10";



/*2.For Loop
Longhand : 
*/
const fruits = ['mango', 'peach', 'banana'];
for (let i = 0; i < fruits.length; i++);

//Shorthand :
for (let fruit of fruits);
//jika kamu hanya mau akses index :
for (let index in fruits);
//ini juga bekerja jika kamu mau access key sebagai literal object:
const obj = {continent: 'Africa', country: 'Kenya', city: 'Nairobi'}
for (let key in obj)
  console.log(key) // output: continent, country, city



/*3.Arrow Function
Longhand : 
*/
function sayHello(name) {
    console.log('Hello', name);
  }
  
  setTimeout(function() {
    console.log('Loaded')
  }, 2000);
  
  list.forEach(function(item) {
    console.log(item);
  });

  //Shorthand :
  sayHello = name => console.log('Hello', name);

setTimeout(() => console.log('Loaded'), 2000);

list.forEach(item => console.log(item));



/*4.Template Literals
Longhand : 
*/
const welcome = 'You have logged in as ' + first + ' ' + last + '.'

const db = 'http://' + host + ':' + port + '/' + database;

//Shorthand :
const welcome = `You have logged in as ${first} ${last}`;

const db = `http://${host}:${port}/${database}`;



/*5.Multi-line String
Longhand : 
*/
const lorem = 'Lorem ipsum dolor sit amet, consectetur\n\t'
    + 'adipisicing elit, sed do eiusmod tempor incididunt\n\t'
    + 'ut labore et dolore magna aliqua. Ut enim ad minim\n\t'
    + 'veniam, quis nostrud exercitation ullamco laboris\n\t'
    + 'nisi ut aliquip ex ea commodo consequat. Duis aute\n\t'
    + 'irure dolor in reprehenderit in voluptate velit esse.\n\t'

    //gampang banget cuman pake backtick
//Shorthand :
const lorem = `Lorem ipsum dolor sit amet, consectetur
    adipisicing elit, sed do eiusmod tempor incididunt
    ut labore et dolore magna aliqua. Ut enim ad minim
    veniam, quis nostrud exercitation ullamco laboris
    nisi ut aliquip ex ea commodo consequat. Duis aute
    irure dolor in reprehenderit in voluptate velit esse.`



/*6.Spread Operator
Longhand : 
*/
// joining arrays
const odd = [1, 3, 5];
const nums = [2 ,4 , 6].concat(odd);
// cloning arrays
const arr = [1, 2, 3, 4];
const arr2 = arr.slice()

//Shorthand : 
// joining arrays
const odd = [1, 3, 5 ];
const nums = [2 ,4 , 6, ...odd];
console.log(nums); // [ 2, 4, 6, 1, 3, 5 ]
// cloning arrays
const arr = [1, 2, 3, 4];
const arr2 = [...arr];



/*7.Array.find
Longhand : 
*/
const pets = [
    { type: 'Dog', name: 'Max'},
    { type: 'Cat', name: 'Karl'},
    { type: 'Dog', name: 'Tommy'},
  ]
  
  function findDog(name) {
    for(let i = 0; i<pets.length; ++i) {
      if(pets[i].type === 'Dog' && pets[i].name === name) {
        return pets[i];
      }
    }
  }

  //Shorthand
pet = pets.find(pet => pet.type ==='Dog' && pet.name === 'Tommy');
console.log(pet); // { type: 'Dog', name: 'Tommy' }


//Lanjutin


