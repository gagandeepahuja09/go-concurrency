Rate Limiting Use Cases

* Attacks on the system by malicious users:
    * They could fill up your service's disk.
    * Brute-force access to the resource.
    * DDoS ==> Distributed denial of service.

* Legitimate users can also take the system down
    * If they are performing operations at a high enough volume or the code they are excersizing is buggy.
    * Usually we provide product SLAs. If one user can affect the SLA for all other users, then it would be a terrible design. A user's mental model is that their access to the system is sandboxed and can neither affect nor be affected by other user's activity. If we break that mental model, then it can feel like the system is not well engineered and they may leave.

* Even with one user, rate-limiting is advantageous. A system might be working well under all the common use-cases but for certain scenarios it may behave differently. In distributed systems, this can have a cascading effect on the system.

***************************************************************************************

More advanatages:

* Rate limit allow you to reason about the performance and stability of your system by preventing it from falling outside the boundary you have investigated. This makes the performance a lot more controllable. If we need to expand those boundaries, we can do so in a controlled manner after lots of testing. 

* In scenarios where we are charging for access, this can help maintain a healthy relationship with the client. We can ask them to try the system under heavily constrained limits.
* It protects our users, if the user was adding too much traffic, does the service owner forgive the cost or is the user forced to pay?

***************************************************************************************

Token Bucket Algorithm

* Access token is required and you won't be able to access the resource without it.
* These tokens are stored in a bucket. d ==> Depth of bucket ==> Max. no of tokens it can hold at a time.
* Access ==> remove a token.
* After the limit is reached, you have to queue your request until a token becomes available or deny the request.
* r ==> Rate at which the tokens are added back. ==> Rate limit.
* Using these 2 we can control both the burstiness and the rate limit.
* NOTE: Users may or may not consume the entire buckets in one long stream. The depth only controls the capacity.
* For users that access the system intermittently but want to round-trip as quickly as possible when they do, bursts are nice to have. We just need to ensure that the system can handle all users bursting at once, or that it is statistically improbable that enough users will burst at the same time to affect our system.

***************************************************************************************

Multi Rate Limiter
* We'll probably want to have multiple limits. Eg. fine grained limits: to limit requests per second and coarse grained limits: to limit request per minute or hour or day.
* If we try to do this in a single rate limiter, we might loose a lot of information around the intent of the rate limiter.
* Makes more sense to keep them separate and then combine them into a single rate limiter.
* For implementing this, we create a separate struct called multiLimiter and an interface called RateLimiter so that multiLimiter can recursively define more multiLimiter instances.
* It will only have 2 methods: Limit, Wait.
* While instantiating the rate limiter itself, we'll sort them by their limits. This will help us in easily implementing the 2 methods: Limit, Wait.

***************************************************************************************

Tiered Rate Limiter
* When we are going to rate limit a system, most likely we'll require to limit more than one thing: Eg. we could be limiting no. of API requests to endpoint A, no. of API requests to endpoint B, rate limit on resources like disk-access, network-access, etc.
* We can reuse the multiLimiter component used above as it can be recursively used to pass more multiLimiters and hence has the option of using both kinds of rate limiting approach together. 