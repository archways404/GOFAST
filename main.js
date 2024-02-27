function main() {
	// Generate a random array of numbers
	const array = generateRandomArray(1000000, 1, 1000000);

	// Measure the start time
	const startTime = performance.now();

	// Find a random number within the array
	const randomNumber = findRandomNumber(array);

	// Measure the end time
	const endTime = performance.now();

	// Calculate the execution time
	const executionTime = endTime - startTime;

	console.log('Random number within the array:', randomNumber);
	console.log('Execution time:', executionTime, 'milliseconds');
}

// generateRandomArray generates a random array of numbers with the specified length and range
function generateRandomArray(length, min, max) {
	const array = [];
	for (let i = 0; i < length; i++) {
		array.push(Math.floor(Math.random() * (max - min + 1)) + min);
	}
	return array;
}

// findRandomNumber finds a random number within the array
function findRandomNumber(array) {
	const index = Math.floor(Math.random() * array.length);
	return array[index];
}

main();
