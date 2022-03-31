When working with concurrent code, there are 2 options for safe operations:
    * Synchronization primitives for sharing memory. eg. sync.Mutex
    * Synchronization via communication eg. channels

Options which are implicitly safe within multiple concurrent operations:
* Immutable data.
    * Create new copy of the data and modify rather than directly modifying it.
    * Apart from ease to dev, it also helps in faster programs as reduces the size or eliminates the critical section.
* Data protected by confinement.

* Idea of confinement: Ensuring information is only ever available from one concurrent process. 2 Kinds of confinement: Ad-hoc, Lexical.

* Ad-hoc: Convention set by a community, hence high chances of anything going wrong. We'll need to do some static-analytics on our code every time someone commits some code. This is very difficult to manage within teams.

* Lexical: Compiler enforces the confinement. Eg. read-only and write-only channels.

* Channels are an easier example since they are concurrent-safe. Next example is for bytes.Buffer, a non concurrent-safe data structure.
* Because of the lexical scope, we have made it impossible to do the wrong thing.

* Why pursue confinement if we have synchronization available to us?
    * Improved performance and reduced cognitive load on devs.
    * Concurrent code which utilizes lexical confinement has the benefit of being simpler to understand.