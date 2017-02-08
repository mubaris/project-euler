#include <stdio.h>

int reversed(int n) {
	int reversed = 0;
	for( ; n > 0; n = n/10) {
		reversed = 10 * reversed + n%10;
	}
	return reversed;
}

int isPalindrome(int n) {
	return (int) n == reversed(n);
}

int main() {
	int largestPalindrome = 0;
	int a = 100;
	for( ; a <= 999; a++) {
		int b = 100;
		for( ; b <= 999; b++) {
			int c = a * b;
			if(isPalindrome(c) && c > largestPalindrome) {
				largestPalindrome = c;
			}
		}
	}
	printf("%d\n", largestPalindrome);
}

