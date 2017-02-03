var sum = 0;

function fibonacci(i) {
	if(i == 1) {
		return 1;
	} 
	else if(i == 2) {
		return 2;
	}
	else {
		return fibonacci(i - 1) + fibonacci(i - 2);
	}
}

for(i = 2; fibonacci(i) <= 4000000; i = i + 3) {
	sum += fibonacci(i);
}

console.log(sum)
