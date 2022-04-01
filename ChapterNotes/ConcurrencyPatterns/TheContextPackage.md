* In concurrent programs, it's necessary to preempt operations because of timeouts, cancellations, or failures of another portion of the system.

* The done channel is somewhat limited. Would be better convey more information like: why the cancellation was occuring or whether or not our function has a deadline by which it needs to complete.

2 Key Points about the structure of Context interface:
* There is nothing present that can mutate the state of the underlying structure.
* There is nothing that allows the function accepting the Context to cancel it. This protects function up the call-stack from children cancelling the context.

* If a context is immutable, how do we affect the behaviour of cancellations in functions below a current function in the call-stack? This is done by all these package level functions:
    func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
    func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
    func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)

* All of these functions take in a context and return one as well.

* WithCancel returns a new context which closes its done channel when the returned cancel function is called.
* WithDeadline returns a new context which closes its done channel when the machine's clock advances past the given deadline.

* If a function needs to cancel functions below it's call graph, it will call one of these functions and pass in the Context it was given, and then pass the context returned into its children. 
* Else just directly pass the parent context. 

* Instances of Context may look similar but may change at every stack-frame.
* For this reason, it's important to always pass instances of Context into our functions.