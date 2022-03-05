* The potential performance gains from implementing the solution to a problem in a parallel manner are bounded by how much of the program must be written in a sequential manner.
* Eg. GUI based app is waiting for user input.(Human interaction). No. of cores won't matter.
* Calculating digits of pi. ==> Embarrasingly parallel. Take instances of your program and run it on more CPUs or machines.  
* Scaling horizontally became much more popular in early 2000s with cloud computing.
* Machines became ephemeral. Solutions could span multiple machines and even global regions.
* Cloud computing made those solutions possible, which were earlier only possible by tech giants.
* Challenges with cloud computing:
    * Provisioning these resources.
    * Communicating between them.(Consul)
    * Aggregating and storing the result.
    * Most difficult: Figuring out how to model code concurrently.
* It enabled all kinds of properties like:
    * Rolling upgrades.
    * Elastic horizontally scalable architecture.
    * Geographic Distribution

********************************************************************************

Why Is Concurrency Hard?

1. Race Conditions
* This occurs when two or more operations must execute in certain order but the program has not been written in a way to guarantee this.
* Most of the data races occur because the developer is thinking about program sequentially. They assume that if one line occurs before the other, it must execute first.
* Bad idea: Adding time.sleep to make parallel code predictable.
    * We have made our program inefficient + We have made our program probabilistic.

2. Atomicity
* Within the "context" that it is operating, it is indivisible or uninterruptable.
* "Context" is important here because something may be atomic in one context but not in another.
* Eg: Atomic within the context of the process may not be atomic in the context of OS.
* Context could be: os, application, process, DB, etc specific.
* FUN FACT: One company created a bot to automatically play a game. As per anti-cheating check a program would run a check by scaning the memory of the host machine.
* Scanning the memory on the machine is considered as an atomic operation. Hardware interrupts were added to hide itself before this scanning started.
    * Retrieve the value of i
    * Increment i
    * Store the value of i
* While all 3 are atomic in an application context. When combined, they won't necessarily be atomic.
* It the context is a program with no concurrent processes, then it will be concurrent.
* It the context is a goroutine which doesn't expose i to other goroutines, then it is atomic.

* Why is atomicity important? If something is atomic, then it is safe with concurrent contexts.
* It can even serve as way to optimize concurrent programs.
* Most statement aren't atomic. Let alone functions, programs. How do we solve this? We can force atomicity by employing various techniques.
* The art then becomes which area of your code need atomicity and at what granularity.