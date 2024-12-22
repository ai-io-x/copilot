To find and print prime numbers under one billion in C, you can use the Sieve of Eratosthenes algorithm. Below is an example implementation in C:

```c


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
```

This code defines a `sieve_of_eratosthenes` function that creates a boolean array to mark prime numbers and iterates over the array to mark non-prime numbers. The `main` function sets the limit to one billion and calls the `sieve_of_eratosthenes` function to print all primes under one billion.

__How to Use GCC to Compile a C Program on Linux__: `sudo apt install build-essential` and `gcc -o app app.c`