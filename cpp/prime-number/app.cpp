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