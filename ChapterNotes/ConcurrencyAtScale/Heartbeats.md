* A way for concurrent processes to signal life to outside parties.
    * They allow us insights into our system.
    * They make testing the system deterministic when otherwise it might not be.

* 2 diff types of heartbeats:
    * Occur on a time interval.
    * Occur at the beginning of a unit of work.

* Our goroutine might be sitting around for a while waiting for something to happen. A heartbeat is a way to signal to its listeners that everything is well & silence is expected.

* In a properly functioning system, heartbeats aren't that interesting. We might use them to gather statistics regarding idle time, but the utility of heartbeats really shines when our goroutine isn't working as expected.

* By using heartbeats, we can avoid deadlocks. It also makes the logic more deterministic by allowing us to put timeouts.

* Adding a buffer of one to the heartbeat channel will ensure that there is always one pulse sent out even if no one is listening in time for the send to occur.

* Why do we set up a separate select block for heartbeat?