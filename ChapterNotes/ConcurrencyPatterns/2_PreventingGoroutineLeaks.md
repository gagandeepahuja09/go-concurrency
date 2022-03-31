* Regardless of how small our goroutine is, we don't want to leave them lying about our process. Goroutine has a few paths to termination: 
    * When it has completed it's work.