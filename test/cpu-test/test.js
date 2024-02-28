function fibonacci(n) {
	if (n <= 1) return n;
	return fibonacci(n - 1) + fibonacci(n - 2);
}

console.time('fibonacci');
console.log(fibonacci(43)); // Adjust the input for deeper performance analysis
console.timeEnd('fibonacci');
