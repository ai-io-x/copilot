To find and print prime numbers under one billion in C++, you can use the Sieve of Eratosthenes algorithm. Below is an example implementation in C++:

```cpp


#include <iostream>
#include <vector>

void sieve_of_eratosthenes(long long limit) {
    std::vector<bool> prime(limit + 1, true);

    for (long long p = 2; p * p <= limit; p++) {
        if (prime[p] == true) {
            for (long long i = p * p; i <= limit; i += p) {
                prime[i] = false;
            }
        }
    }

    for (long long p = 2; p <= limit; p++) {
        if (prime[p]) {
            std::cout << p << std::endl;
        }
    }
}

int main() {
    long long limit = 1000000000LL;
    sieve_of_eratosthenes(limit);
    return 0;
}
```

This code defines a `sieve_of_eratosthenes` function that creates a boolean vector to mark prime numbers and iterates over the vector to mark non-prime numbers. The `main` function sets the limit to one billion and calls the `sieve_of_eratosthenes` function to print all primes under one billion.

__How to Use GCC to Compile a C Program on Linux__: `sudo apt install build-essential` and `g++ -o app app.cpp`