# One-Pass Hash table

#### 
    The One-Pass Hash table approach processes input data in a single sequential scan (one pass).
    The Key is using an in-memory hash table (dictionary) to instantly store and look up values encountered so far. This allows operations that might normally require a second pass(like finding a complement or checking a frequency) to happen instantaneously within the same loop.
    Memorization point: Read once, use hash table to remember previous values instantly. Time complexity is highly efficient at O(N)
####