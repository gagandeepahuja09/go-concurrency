* A way for concurrent processes to signal life to outside parties.
    * They allow us insights into our system.
    * They make testing the system deterministic when otherwise it might not be.

* 2 diff types of heartbeats:
    * Occur on a time interval.
    * Occur at the beginning of a unit of work.

* Our goroutine might be sitting around for a while waiting for something to happen. A heartbeat is a way to signal to its listeners that everything is well & silence is expected.

* In a properly functioning system, heartbeats aren't that interesting. We might use them to gather statistics regarding idle time, but the utility of heartbeats really shines when our goroutine isn't working as expected.