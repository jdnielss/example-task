const number = Number(prompt('Enter Number : '));

if (number > 0) {
  console.log(`${number} is positive`);
} else if (number < 0) {
  console.log(`${number} is negative`);
} else {
  console.log(`${number} is zero`);
}
