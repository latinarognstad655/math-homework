def fibonacci_sequence(n):
    """
    Generate a Fibonacci sequence up to n.
    
    Parameters:
    n (int): The length of the Fibonacci sequence to generate.
    
    Returns:
    list: A list containing the first n numbers in the Fibonacci sequence.
    """
    sequence = []
    a, b = 0, 1
    while len(sequence) < n:
        sequence.append(a)
        a, b = b, a + b
    return sequence

# Example usage
n = 10
fibonacci_sequence(n)
