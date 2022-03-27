Key Requirement for Channels:

* All the execution stages must accept and return the same type.(Same function parameters and same return type)
* Using channels, we have the advantage of executing each stage of the pipeline concurrently.
* Generator function converts a discrete set of values into a stream of data on channel.

* All the execution stages as well as the generator run concurrently(each have a go routine and a for select inside the goroutine) so that they don't have to wait for one execution stage to complete in order for the other to start.

* done channel is used in each method for prevent goroutine leaks.

***************************************************************************************

Repeat Take Pattern in Pipelines
* Repeat take both are generators.
* Repeat will infinitely keep on sending the values of the interface in a cycle.
* Take will only ready the first num values from the channel.

* It is OK to use interface{} as types for channels so that we can use them as a standard library of pipeline patterns.

* When we want to deal with specific type, we can add stages with type assertions. The performance overhead for adding an extra stage is negligible.