import time
import cProfile


def main():
    start_time = time.time()  # Get the current time
    n = 0
    for i in range(100000000):
        n += 1
    end_time = time.time()  # Get the current time again
    elapsed_time = (
        end_time - start_time
    )  # Subtract start from end to get the elapsed time
    print("With main function took :", elapsed_time, "seconds")


cProfile.run("main()")

# With main function took : 3.6934521198272705 seconds
#          7 function calls in 3.693 seconds

#    Ordered by: standard name

#    ncalls  tottime  percall  cumtime  percall filename:lineno(function)
#         1    0.000    0.000    3.693    3.693 <string>:1(<module>)
#         1    3.693    3.693    3.693    3.693 main_function.py:5(main)
#         1    0.000    0.000    3.693    3.693 {built-in method builtins.exec}
#         1    0.000    0.000    0.000    0.000 {built-in method builtins.print}
#         2    0.000    0.000    0.000    0.000 {built-in method time.time}
#         1    0.000    0.000    0.000    0.000 {method 'disable' of '_lsprof.Profiler' objects}
