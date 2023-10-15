import cProfile

cProfile.run(
    """
import time

start_time = time.time()
n = 0
for i in range(100000000):
    n += 1
end_time = time.time()
elapsed_time = end_time - start_time
print("Without main function took :", elapsed_time, "seconds")
"""
)


# Without main function took : 6.110252857208252 seconds
#          6 function calls in 6.110 seconds

#    Ordered by: standard name

#    ncalls  tottime  percall  cumtime  percall filename:lineno(function)
#         1    6.110    6.110    6.110    6.110 <string>:1(<module>)
#         1    0.000    0.000    6.110    6.110 {built-in method builtins.exec}
#         1    0.000    0.000    0.000    0.000 {built-in method builtins.print}
#         2    0.000    0.000    0.000    0.000 {built-in method time.time}
#         1    0.000    0.000    0.000    0.000 {method 'disable' of '_lsprof.Profiler' objects}
