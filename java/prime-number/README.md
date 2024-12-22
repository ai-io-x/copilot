To find and print prime numbers under one billion in Java, you can use the Sieve of Eratosthenes algorithm. This algorithm is efficient for generating a list of prime numbers up to a given limit. Below is an example implementation in Java:

# Explanation
__sieveOfEratosthenes Method__:

Initializes a boolean array `prime` where each index represents whether the number is prime.
Iterates over the array and marks the non-prime numbers.
Collects and returns all prime numbers.

__main Method__:

Sets the limit to one billion.
Calls the `sieveOfEratosthenes` method to get all primes under one billion.
Prints each prime number.

# Considerations
Time and Space Complexity: This algorithm is efficient, but generating and storing prime numbers up to one billion can require significant time and memory. Ensure your system has enough resources.
__Output Size__: Printing all prime numbers under one billion will generate a very large output. Consider writing the output to a file or limiting the number of primes you print if you are running this in a typical environment.
Real-World Use Case
The Sieve of Eratosthenes is commonly used in various fields such as cryptography, number theory, and computer science research, where prime number generation is essential.

# Exception in thread "main" java.lang.OutOfMemoryError: Java heap space

__How can I find Java heap size and memory used__: `java -XX:+PrintFlagsFinal -version | grep HeapSize` or `jps -lvm`

__How to increase memory allocated to java__: with `-Xmx1024m` like `java -Xmx3072m PrimeNumbers`