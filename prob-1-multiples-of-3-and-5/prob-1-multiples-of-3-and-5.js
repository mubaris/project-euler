var total = 0;

/*for(var i = 0; i < 1000; i++) {
	if((i % 3 == 0) || (i % 5 == 0)) {
		total += i;
	}
}*/

//Optimized Solution

function sumDivisibleBy(n) {
	var p = parseInt(999 / n, 10)
	return parseInt(n * p * (p + 1) / 2, 10)
}

total = sumDivisibleBy(3) + sumDivisibleBy(5) - sumDivisibleBy(15)

console.log(total)
