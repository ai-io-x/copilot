def sieve_of_eratosthenes(limit):
    # Create a boolean array "prime[0..limit]" and initialize all entries as True.
    # A value in prime[i] will be False if i is not a prime, else True.
    prime = [True for _ in range(limit + 1)]
    p = 2
    while p * p <= limit:
        # If prime[p] is still True, then it is a prime
        if prime[p]:
            # Updating all multiples of p to not prime
            for i in range(p * p, limit + 1, p):
                prime[i] = False
        p += 1

    # Collect all prime numbers
    prime_numbers = []
    for p in range(2, limit + 1):
        if prime[p]:
            prime_numbers.append(p)
    return prime_numbers

def main():
    limit = 10**9  # One billion
    primes = sieve_of_eratosthenes(limit)

    # Printing all primes under one billion
    # Note: This might take a considerable amount of time and memory
    for prime in primes:
        print(prime)

if __name__ == "__main__":
    main()