* In distributed systems, it's both easy for something to go wrong in our system and difficult to understand why it happened.

* We discussed how to propagate errors but not on what these errors should look like and how they should flow through the system.

* Error indicates that the system reached a state in which it cannot fulfil the request made by the user either explicitly or implicitly. 

* It should contain the following info:

* What happened: Error message. eg "disk full", "socket closed".
* When and where it occurred:
    * Complete stack trace. It shouldn't be part of the error message but should be easily accessible.
    * It should contain info regarding the context it's running within. ie machine eg. stack, dev, prod.
    * Time when it occurred in UTC.
