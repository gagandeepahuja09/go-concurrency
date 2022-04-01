* As we develop our programs, we should give our error paths the same attention we give our algorithms.

* Most fundamental question: Who should be responsible for handling the error? At some point, the program needs to stop ferrying the error up the stack trace and do something with it.

* With concurrent programs the question becomes more complex since a concurrent process is operating independent to its parent or siblings.
* In the naive example we have put the goroutine in an awkward position and assumes that the main method is paying attention to that.
* Rather than this, the error should be returned as part of the channel struct.