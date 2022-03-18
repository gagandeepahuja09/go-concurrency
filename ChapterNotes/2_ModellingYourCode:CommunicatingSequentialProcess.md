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

What Is CSP?
* In this paper on CSP, Hoare suggests that input & output are 2 overlooked primitives of programming - particularly concurrent code.
* His programming language contained primitives to model input and output, or communication between processes correctly.
* ! for sending input to a process.
* ? for reading input out a process.
* It's similarities with Go channel are very apparent.

* cardreader?cardimage ==> From cardreader read a card and assign it's value to the varialbe cardimage.
* lineprinter!lineimage ==> To lineprinter send the value of lineimage for printing.

* The above 2 message passing a very similar to "reading from" and "writing to" a channel.

* X?(x, y) ==> From a process named X input 2 values and assign them to variables x,y.
* DIV!(3*a + b, 12) ==> To process DIV output the 2 specified values.

* *[c:character; west?c -> east!c] ==> This reads all characters output by west and outputs them one by one to east. The repetition terminates when the process west terminates.

* The above examples indicate how a language with first-class support for modelling communication makes solving problems simpler and easier to comprehend.
* Most languages favor sharing and synchronizing access to the memory as compared to CSP's message passing style.
* Memory sharing, synchronizing isn't bad. Sometime it is appropriate, even in Go.
* However the shared memory can be difficult to utilize correctly - especially in large or complicated programs.

********************************************************************************

How does this help you?

* The comparison b/w go and other languages is generally a comparison between a goroutine & a thread or a channel & a mutex.
* Goroutines free us from having to think about our problem in terms of parallelism.
* Example: building a web server. Setting aside language frameworks, in a language which only offers thread abstraction, we would have to think about the following points: 
    * Does my language naturally support threads or will I have to pick a library?
    * How heavy are threads in the operating system?
    * How does the OS where my program will be running in handle threads differently?
    * I should create a worker pool to constrain the number of threads I create. How do I find the optimal number?
* All of these are important to consider but none of them directly concern the problem we're trying to solve.
* All this is acheived in goroutines by a promise that Go makes us, that goroutines are lightweight, and we normally won't have to worry about creating one.
* There a appropriate times to consider the # of running goroutines but doing this upfront is premature optimisation. In constrast to threads, where we would need to think about all this upfront.

* This mapping also has a side-effect benefit for us. Go's runtime multiplexes goroutines onto OS threads for us and manages their scheduling for us. This means that optimizations to runtime can be made without having to change how we've modeled our problem.
* As advancements in parallelism are made, Go's runtime will improve, as will the performance of your program, all for free.

* Advantage 2: Since Go's runtime is managing the scheduling of goroutines for you, it can introspect on things like goroutines blocked waiting for I/O and intelligently reallocate OS threads to goroutines which are not blocked. This also increases the performance of our code.

* Advantage 3: Because the problems we work on as developers are naturally concurrent more often than not, we'll be writing concurrent code at a finer level of granularity than we perhaps would in other languages. Eg. goroutine for every user, instead of multiplexed into a thread pool. This level of granularity helps to dynamically scale the amount of parallelism - Amdahl's law in action.
