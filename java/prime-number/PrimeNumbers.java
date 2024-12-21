import java.util.ArrayList;
import java.util.List;

public class PrimeNumbers {

    public static List<Integer> sieveOfEratosthenes(int limit) {
        // Create a boolean array "prime[0..limit]" and initialize all entries as true.
        // A value in prime[i] will be false if i is not a prime, else true.
        boolean[] prime = new boolean[limit + 1];
        for (int i = 0; i <= limit; i++) {
            prime[i] = true;
        }

        for (int p = 2; p * p <= limit; p++) {
            // If prime[p] is still true, then it is a prime
            if (prime[p]) {
                // Updating all multiples of p to not prime
                for (int i = p * p; i <= limit; i += p) {
                    prime[i] = false;
                }
            }
        }

        // Collect all prime numbers
        List<Integer> primeNumbers = new ArrayList<>();
        for (int p = 2; p <= limit; p++) {
            if (prime[p]) {
                primeNumbers.add(p);
            }
        }
        return primeNumbers;
    }

    public static void main(String[] args) {
        int limit = 1000000000; // One billion
        List<Integer> primes = sieveOfEratosthenes(limit);

        // Printing all primes under one billion
        // Note: This might take a considerable amount of time and memory
        for (int prime : primes) {
            System.out.println(prime);
        }
    }
}