* Regardless of how small our goroutine is, we don't want to leave them lying about our process. Goroutine has a few paths to termination: 
    * When it has completed it's work.
    * When it cannot continue it's work due to an unrecoverable error.
    * When it's told to stop working.

* The 3rd part is of most interest to us. The parent goroutine(generally the main goroutine) should be able to tell all of its children goroutines to terminate.

* At the part where we are reading from the channel returned by the goroutine to the main goroutine, we call it the join part as we are joining the spawned goroutine and main goroutine.

* The goroutine which is not returned will remain in memory for the lifetime of the process.
* We will even deadlock if the spawned goroutine is an infinite loop.
* In the worst case a long-running main goroutine could continue to spin up goroutines throughout its lifetime causing memory utilization to increase.

* If we just join or read from a channel without using done, what is meant to be run in the defer statement won't run.

* If a goroutine is responsible for creating a goroutine, then it should also be response for ensuring that it can stop the goroutine.