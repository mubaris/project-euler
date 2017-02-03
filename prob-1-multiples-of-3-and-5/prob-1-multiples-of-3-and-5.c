#include <stdio.h>

int main() {
	int total = 0;
	//int i = 0;

	/*for(i = 0; i < 1000; i++) {
		if((i % 3 == 0) || (i % 5 == 0)) {
			total += i;
		}
	}*/

	//Optimized Solution
	
	int target = 999;

	int sumDivisibleBy(int n) {
		int p;
		p = target / n;
		return (n * p * (p + 1)) / 2;
	}

	total = sumDivisibleBy(3) + sumDivisibleBy(5) - sumDivisibleBy(15);
	
	printf("%d\n", total);
}
