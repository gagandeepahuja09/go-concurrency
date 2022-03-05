* The potential performance gains from implementing the solution to a problem in a parallel manner are bounded by how much of the program must be written in a sequential manner.
* Eg. GUI based app is waiting for user input.(Human interaction). No. of cores won't matter.
* Calculating digits of pi. ==> Embarrasingly parallel. Take instances of your program and run it on more CPUs or machines.  
* Scaling horizontally became much more popular in early 2000s with cloud computing.
* Machines became ephemeral. Solutions could span multiple machines and even global regions.
* Cloud computing made those solutions possible, which were earlier only possible by tech giants.
* Challenges with cloud computing:
    * Provisioning these resources.
    * Communicating between them.(Consul)
    * Aggregating and storing the result.
    * Most difficult: Figuring out how to model code concurrently.
* It enabled all kinds of properties like:
    * Rolling upgrades.
    * Elastic horizontally scalable architecture.
    * Geographic Distribution