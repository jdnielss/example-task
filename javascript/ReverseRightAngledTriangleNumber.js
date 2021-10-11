// Create functions to find the reverse right angled triangle number of a given number.
const triangle = (input) => {
  // Validate the input.
  if (typeof input === "number") {
    // Loop through the input.
    for (i = input; i > 0; i--) {
      for (j = 1; j <= i; j++) {
        // Print the number.
        process.stdout.write("" + j);
      }
      // Add a new line.
      console.log();
    }
  } else {
    // Output if the input is not a number.
    console.log("Input must be number");
  }
};
// Call the function.
triangle(8);
