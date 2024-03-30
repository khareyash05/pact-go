# Contract testing 
Contract testing is an essential technique for ensuring reliable communication between distributed systems. It involves testing the interactions between the components of a system by defining a set of agreed-upon contracts that specify the expected behavior of each component. This technique helps to catch any breaking changes that may occur during the development process, reducing the risk of system failure in production.
![1_64wLFdRJG6xjxZhSsrsCCQ](https://github.com/khareyash05/pact-go/assets/60147732/78027a2f-6e02-4521-a49c-13b4bb055098)
We in general generate integration tests for our applications to check whether microservices goes hand in hand.<br>
Most famous integration testing is the e2e testing which test all microservices as whole as they will be when in production. 

## Problems with e2e integration testing
1. Time Consuming
2. Not easy to detect faults
3. Not compatible with microservices changing demands as it might cause other microservices to break.
<br>Thus Contract Testing came into play

## What is Pact?
1. Pact is a popular contract testing tool that provides a framework for implementing contract testing in a variety of languages and frameworks, including Go<br>
2. Pact works by generating mock servers that simulate the behavior of a provider system, allowing the consumer system to test its interactions with the provider and generate a contract that specifies the expected behavior of the provider
   


## Steps to perform Contract Testing

#1 Generate/update a contract:
1. cd /pkg/client
2. go test

#2 Run provider test to validate it against the contract:
1. cd /pkg/server
2. go test

## Output
![1_L_UfSS7tKnAxe5lo2cENZw](https://github.com/khareyash05/pact-go/assets/60147732/53e4ff39-b0f1-4bce-b33e-454a1745bad2)<br>
Testing and validating the contract

![1_FajCLxCUA1fZEYyoP_s9zg](https://github.com/khareyash05/pact-go/assets/60147732/e4a842be-850a-49c3-b75b-ec15658878c1)
