fn sieve_of_eratosthenes(limit: usize) -> Vec<usize> {
    // Create a boolean array "prime[0..limit]" and initialize all entries as true.
    // A value in prime[i] will be false if i is not a prime, else true.
    let mut prime = vec![true; limit + 1];
    let mut p = 2;

    while p * p <= limit {
        // If prime[p] is still true, then it is a prime
        if prime[p] {
            // Updating all multiples of p to not prime
            let mut i = p * p;
            while i <= limit {
                prime[i] = false;
                i += p;
            }
        }
        p += 1;
    }

    // Collect all prime numbers
    let mut prime_numbers = Vec::new();
    for p in 2..=limit {
        if prime[p] {
            prime_numbers.push(p);
        }
    }

    prime_numbers
}

fn main() {
    let limit = 1_000_000_000; // One billion
    let primes = sieve_of_eratosthenes(limit);

    // Printing all primes under one billion
    // Note: This might take a considerable amount of time and memory
    for prime in primes {
        println!("{}", prime);
    }
}