#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

void sieve_of_eratosthenes(long long limit) {
    bool *prime = malloc((limit + 1) * sizeof(bool));
    for (long long i = 0; i <= limit; i++) {
        prime[i] = true;
    }

    for (long long p = 2; p * p <= limit; p++) {
        if (prime[p] == true) {
            for (long long i = p * p; i <= limit; i += p) {
                prime[i] = false;
            }
        }
    }

    for (long long p = 2; p <= limit; p++) {
        if (prime[p]) {
            printf("%lld\n", p);
        }
    }

    free(prime);
}

int main() {
    long long limit = 1000000000LL;
    sieve_of_eratosthenes(limit);
    return 0;
}