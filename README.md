# golang - tutorials and helper things
few simple resources on concurrency taks and problems working with golang. 

## Concurrency 
Thinking about large programs made up of many smaller components, how do we run all these at the same time?? Use Concurrency.
e.g. think of a web server 

### How? - GoRoutines and Channels

GoRoutine * func which can run concurrently with other funcs
          * returns immediately to next line in ord of ops

Channel   * provide way for two goroutines to talk to each other
