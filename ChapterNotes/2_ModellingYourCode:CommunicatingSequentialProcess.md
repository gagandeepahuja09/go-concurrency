* In this chapter we'll discuss on why go got so much right.

********************************************************************************

Difference between concurrency and parallelism
* CONCURRENCY is a property of the code. While PARALLELISM is a property of the running program.
* If I write my program with the intention that the 2 chunks of the program run in parallel, then is it actually guaranteed? What happens if there is only core? They'll run sequentially with context switches. Over a coarse enough granularity, they appear to be running in parallel.
* Few important points:
    1. We don't write parallel code. Only concurrent code hoping to run in parallel.
    2. It is possible-maybe even desirable to be ignorant of whether our concurrent code is running in parallel.
        This is made possible only by the layer of abstractions that lay beneath our program's model: 
            1. the concurrency primitives.
            2. the program's runtime.
            3. the operating system.
            4. the platform the operating system runs on.(could be containers/VMs)
            5. CPUs.
        These platforms allow us to make distinction b/w concurrency and parallelism and in turn which gives us the power and flexibility to express ourselves.
    3. Parallelism is a function of time or context. Eg. context = 5 seconds, switchTime = 1s ==> running in parallel. context = 1s, switchTime = 1s ==> not parallel.
    
    * This was about time. Context could be other context as well like machine or process or thread. 
    * We can safely assume that two logics running on a different machine are isolated from each other(machine context). We can also assume at a process level that 2 calculator processes running on the same machine are isolated from each other.
    * What about threads? We can face all problems discussed earlier: race condition, deadlocks, livelocks, starvation.
    * Unfortunately most concurrent logic in our industry is written at one of the highest level of abstraction: OS threads.
    * Before Go was released, this was where the chain of abstraction ended.
    * For concurrent code, we had to model our program in terms of threads and synchronize the access to the memory between them.
    * If we had a lot of things that we wanted to model concurrently and our machine couldn't handle that many threads, we created a thread pool. And multiplexed our operations onto the thread pool.
    * With goroutines, threads are also there but with the extra layer of abstraction, we havely to rarely think about our problem in terms of OS threads. Instead we model things in goroutines and channels, and ocassionally shared memory.

********************************************************************************