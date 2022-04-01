* Combining one or more done channels into a single done channel which closes if any of its component closes.

* We can do this via select statement also, but we might not know the number of done channels we are working with at runtime.

* Or-channel: Creates composite done channel through recursion and goroutine.