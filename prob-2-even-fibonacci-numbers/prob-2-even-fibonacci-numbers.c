#include <stdio.h>

int fibonacci(int n) {
	if(n == 1) {
		return 1;
	}
	else if(n == 2) {
		return 2;
	}
	else {
		return fibonacci(n - 1) + fibonacci(n - 2);
	}
}

int main() {
	int i = 0;
	int sum = 0;
	for(i = 2; fibonacci(i) <= 4000000; i = i + 3) {
		sum += fibonacci(i);
	}
	printf("%d\n", sum);
}
