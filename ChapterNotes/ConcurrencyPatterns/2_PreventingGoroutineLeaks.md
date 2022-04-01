* Regardless of how small our goroutine is, we don't want to leave them lying about our process. Goroutine has a few paths to termination: 
    * When it has completed it's work.
    * When it cannot continue it's work due to an unrecoverable error.
    * When it's told to stop working.

* The 3rd part is of most interest to us. The parent goroutine(generally the main goroutine) should be able to tell all of its children goroutines to terminate.