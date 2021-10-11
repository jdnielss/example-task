// Declare a variable and assign it a value.
const math = 68;
const art = 69;
const history = "69";
const science = 70;
const average = (math + art + history + science) / 4;

// Validation for the average grade
if (math === NaN && art === NaN && history === NaN && science === NaN) {
  // Conditionals to determine the grade
  if (average < 60) {
    grade = "E";
  } else if (average < 70) {
    grade = "D";
  } else if (average < 80) {
    grade = "C";
  } else if (average < 90) {
    grade = "B";
  } else {
    grade = "A";
  }
  // Output
  console.log("Average: " + average);
  console.log("Grade: " + grade);
} else {
  // Output if the user enters a non-number
  console.log("Input must be number");
}
