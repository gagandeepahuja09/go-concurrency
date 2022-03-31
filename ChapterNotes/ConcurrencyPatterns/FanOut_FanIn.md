* We often have the ability to reuse the stages of the pipeline multiple times.
* We can use this to parallelize the pulls from an upstream stage.

* Fan-out: Starting multiple goroutines to handle input from the pipeline.
* Fan-in: Combining multiple results into one channel.

* Both properties should hold in order to apply fan-out pattern: 
    * The stage doesn't rely on the values that it had calculate before.
    * The stage must be long running.

* Order-independence is important because we have no guarantee on what order concurrent copies of our stage will run nor what order they will return.

* The process of fanning out a stage in a pipeline is extraordinarily easy.
* All we have to do is start multiple versions of that stage.
    ==> make([]chan <-int, numFinders)
* In example we are starting as many copies of this stage as we have CPUs. For production empirical testing would be required to pick the optimal number.


* Fan In
    * Take a variadic parameter with all channels as input(read only).
    * Return the final channel(read only).
    * Range through the channels. Run that many goroutines.
    * Each goroutine will read all values from that channel and output to the output multiplexed channel.
    * Use waitgroups to ensure that we wait for all goroutines to finish before returning the final output channel.