# Distributed System vs Sequential Processing: Performance Comparison

## Distributed System (10 Workers)

### Test Case 1: Sum of Numbers from 1 to 1,000,000
- **Total Sum**: 500000500000
- **Time Taken**: 1.089606ms

### Test Case 2: Sum of Numbers from 1 to 100,000,000
- **Total Sum**: 5000000050000000
- **Time Taken**: 10.251562ms

---

## Sequential Processing (Single Worker)

### Test Case 1: Sum of Numbers from 1 to 1,000,000
- **Total Sum**: 500000500000
- **Time Taken**: 301.131µs

### Test Case 2: Sum of Numbers from 1 to 100,000,000
- **Total Sum**: 5000000050000000
- **Time Taken**: 24.27606ms

---

## Observations

- For smaller workloads (e.g., summing numbers from 1 to 1,000,000), the sequential approach is faster due to the overhead of distributed communication.
- For larger workloads (e.g., summing numbers from 1 to 100,000,000), the distributed system significantly outperforms the sequential approach by dividing the task among multiple workers and processing in parallel.