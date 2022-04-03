Rate Limiting Use Cases

* Attacks on the system by malicious users:
    * They could fill up your service's disk.
    * Brute-force access to the resource.
    * DDoS ==> Distributed denial of service.

* Legitimate users can also take the system down
    * If they are performing operations at a high enough volume or the code they are excersizing is buggy.
    * Usually we provide product SLAs. If one user can affect the SLA for all other users, then it would be a terrible design. A user's mental model is that their access to the system is sandboxed and can neither affect nor be affected by other user's activity. If we break that mental model, then it can feel like the system is not well engineered and they may leave.

* Even with one user, rate-limiting is advantageous. A system might be working well under all the common use-cases but for certain scenarios it may behave differently. In distributed systems, this can have a cascading effect on the system.