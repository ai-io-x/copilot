package main

import (
    "fmt"
    "math"
)

func sieveOfEratosthenes(limit int) []int {
    // Create a boolean array "prime[0..limit]" and initialize
    // all entries it as true. A value in prime[i] will
    // finally be false if i is Not a prime, else true.
    prime := make([]bool, limit+1)
    for i := range prime {
        prime[i] = true
    }

    for p := 2; p*p <= limit; p++ {
        // If prime[p] is not changed, then it is a prime
        if prime[p] {
            // Update all multiples of p to not prime
            for i := p * p; i <= limit; i += p {
                prime[i] = false
            }
        }
    }

    // Collect all prime numbers
    var primes []int
    for p := 2; p <= limit; p++ {
        if prime[p] {
            primes = append(primes, p)
        }
    }
    return primes
}

func main() {
    limit := int(math.Pow(10, 9)) // One billion
    primes := sieveOfEratosthenes(limit)

    fmt.Printf("Prime numbers under %d:\n", limit)
    for _, prime := range primes {
        fmt.Println(prime)
    }
}