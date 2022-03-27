* We often have the ability to reuse the stages of the pipeline multiple times.
* We can use this to parallelize the pulls from an upstream stage.

* Fan-out: Starting multiple goroutines to handle input from the pipeline.
* Fan-in: Combining multiple results into one channel.

* Both properties should hold in order to apply fan-out pattern: 
    * The stage doesn't rely on the values that it had calculate before.
    * The stage must be long running.

