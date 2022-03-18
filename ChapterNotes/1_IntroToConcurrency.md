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

1 var data int 
2 go func() {
3    data++
4 }()
5 if data == 0 {
6     fmt.Printf("the value is %v.\n", data)
7 }

* nothing is print if 3 was executed before 5.
* the values is 0 is printed if line 5 and 6 were executed before line 3.
* the values is 1 is printed if line 5 was executed before line 3 but line 3 was executed before line 6.

********************************************************************************

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
* If the context is a program with no concurrent processes, then it will be concurrent.
* If the context is a goroutine which doesn't expose i to other goroutines, then it is atomic.

* Why is atomicity important? If something is atomic, then it is safe with concurrent contexts.
* It can even serve as way to optimize concurrent programs.
* Most statement aren't atomic. Let alone functions, programs. How do we solve this? We can force atomicity by employing various techniques.
* The art then becomes which area of your code need atomicity and at what granularity.

********************************************************************************

3. Memory Access Synchronization
* Critical section: Name for the section of our program that needs exclusive access to a shared resource.
* In race condition example, there are 3 CS: increment, check, print.
* NOTE: Non idiomatic Go code. Not recommended way of doing it.

var memoryAccess sync.Mutex 
var data int 
go func() {
    memoryAccess.Lock()
    data++
    memoryAccess.Unlock()
}()
memoryAccess.Lock()
if data == 0 {
    fmt.Printf("the value is %v.\n", data)
}
memoryAccess.Unlock()

* IMP: We haven't fully solved the issue if we want to execute the code in a certain order. This will only ensure that the 2 blocks have exclusive access but the order will be non-determinisitic. We'll learn later for the tools to solve issues of these kinds of problems.
* We'll later see the performance implications of synchronizing access to memory.

* Two tricky questions which should be answered on the basis of your code:
    * What size should my critical sections be?
    * Are my CS entered and exited repeatedly? 

********************************************************************************

Deadlocks, Livelocks, and Starvation

Deadlock

* Deadlock: All concurrent processes are waiting on one another. In this case, the program will never recover without outside intervention.
* Goroutines can help with deadlock detection and recovery in some cases but not with deadlock prevention.

* In the race_condition/main.go example, the first goroutine has locked a and attempts to lock b. While the 2nd goroutine locks b and attempts to lock a. Both goroutines will wait infinitely on each other.

* Coffman Condition:
* Mutually exclusive: A concurrent process holds exclusive rights to a resource at any point in time.
* Wait for condition: A concurrent process must hold a resource and be waiting on another resource.
* No preemption: A resource held by a process can only be released by that process.
* Circular Wait: A concurrent process(P1) must be waiting on a chain of other concurrent process P2, which are in turn waiting on it(P1). 


* If we can't any of the one condition, we can prevent deadlock.
* In practice, it can be very hard to find.


********************************************************************************

Livelock
* Programs which are actively performing concurrent operations but these programs do nothing to move the state of the program forward. 
* We can also say that it is a scenario where 2 or more concurrent processes are trying to avoid deadlock with no coordination.
* Eg. you are moving in a hallway and someone is coming from the front. You both keep on moving from one side to the other in order to make way for the other. HALLWAY SHUFFLE

* You can go through the following sync packages:
https://pkg.go.dev/sync#Cond => Cond => Condition variable => 
    Broadcast method ==> wakes all goroutines waiting on c(*Cond).
https://pkg.go.dev/sync#Locker ==> interface with Lock and Unlock method. Mutex, RWMutex implement it. 

* Livelocks could be even harder to spot than deadlocks because it can appear that the program is doing work.
* By looking a CPU utilization, you might think that it is doing work.
* Using sync.Cond's Wait method, we ensure that both the process move at same speed.
* Livelocks are a subset of a larger set of problems called "Starvation"

********************************************************************************

Starvation
* When a concurrent process cannot get all the resources it needs to perform work.
* In case of livelock, each resource is starved for a shared lock. Livelock is a special starvation case in which all goroutines are starved equally and no work is accomplished.
* Common scenario: There are greedy processes which are unfairly preventing some of the concurrent processes from accomplishing work as efficiently as possible, or any work at all.
* In the starvation/main.go example, the greedy worker greedily holds on to the lock for the entirity of its work loop, whereas the polite work holds only when it needs to.
* Here the greedy worker, it preventing the polite work to do its work optimally.

* Detecting starvation requires logging metrics.
* This can be detected by determinig if our rate of work is as high as expected.

* Finding a balance: In case of memory access synchronization, we'll have to find a balance preferring coarse-grained synchronization for performance and fine-grained synchronization for fairness.

* Going with fine-grained first is a much better way. We can always broaden the scope.
* Starvation can also come out of go processes. It can come out of CPU, DB connections, file handles, memory, etc: any resource which must be shared is a candidate for starvation.

********************************************************************************

Determining Concurrency Safety

Common questions:
* How to find the right level of abstraction for concurrent code?
* Techniques to create a solution that is both easy to use and modify?
* What is the right level of concurrency for this problem?

* When going through existing code ==> Not always obvious what code is utilizing concurrency and how to utilize the code safely?

    // CalculatePi calculates digits of Pi b/w the begin and end place.
    func CalculatePi(begin, end int64, pi *Pi)

* The above function raise a lot of questions:
    * Am I responsible for instantiating multiple concurrent invocations of this function?
    * All the functions are going to be using the instance of Pi whose address we pass in. Am I responsible for synchronizing the access to this memory of does the function take care of this?

* Comments work wonders in such cases:  
    // CalculatePi calculates digits of Pi b/w the begin and end place.
    // Internally CalculatePi calls FLOOR((end - begin) / 2) concurrent processes 
    // which recursively call CalculatePi. Synchronization to writes are handled internally by the Pi struct.

* It answers important questions like: who is responsible for concurrency and who is responsible for synchronization.
* We can change how we have modeled it. If we take a functional approach and ensure that our function has no side effects, we are removing questions related to synchronization.

    func CalculatePi(begin, end int64) <-chan int

********************************************************************************

Simplicity in the face of complexity
* With Go's concurrency primitives, you can more safely and clearly express your concurrent algorithms.

* How do they make our life easier?
1. Go's Concurrent, Low Latency Garbage Collector:
    * Are they good to have?
    * They prevent work in a problem domain which requires real-time performance or a deterministic performance profile.
    * Pausing all activity in a program to clean up garbage isn't acceptable.
    * Go 1.8 ==> takes only 10-100 microseconds.
    * How does this help us? Memory management is a difficult problem to solve. Coupled with concurrency, it becomes even more difficult.
    * Go ensures that you don't have to forcefully manage memory simply or across concurrent processes.

2. Go's runtime automatically handles multiplexing concurrent operations onto operating system threads.(ABSTRACTION) It allows us to directly map concurrent problems into concurrent constructs instead of starting and managing threads and mapping logic evenly across threads.
* In some languages, we would first need to create a thread pool(colln of threads).
* Then map incoming conns to threads.
* Within each thread loop over all connections to ensure that they all receive some CPU time.
* The connection handling logic needs to be pausible so that it can be shared fairly with other connections.
* The go runtime will handle all this automatically.

3. Makes composing larger problems easy. Channel provides composable, concurrent safe way to communicate between processes.

