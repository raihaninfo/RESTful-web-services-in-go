# RESTful-web-services-in-go

## What is REST architecture?

REST stands for `REpresentational State Transfer.` REST is web standards based architecture and uses HTTP Protocol. It revolves around resource where every component is a resource and a resource is accessed by a common interface using HTTP standard methods. REST was first introduced by `Roy Fielding` in 2000.

In REST architecture, a REST Server simply provides access to resources and REST client accesses and modifies the resources. Here each resource is identified by URIs/ global IDs. REST uses various representation to represent a resource like text, JSON, XML. JSON is the most popular one.

## Characteristics of REST services

### These are the main properties that make REST simple and unique compared to its predecessors:

- **Client-server based architecture:** This architecture is most essential for the modern web to communicate over HTTP. A single client-server may look naive initially, but many hybrid architectures are evolving. We will discuss more of these shortly.

- **Stateless:** This is the most important characteristic of a REST service. A REST HTTP request consists of all the data needed by the server to understand and give back the response. Once a request is served, the server doesn't remember if the request has arrived after a while. So the operation will be a stateless one.

- **Cacheable:** Many developers think a technology stack is blocking their web application or API. But in reality, their architecture is the reason. The database can be a potential tuning piece in a web application. In order to scale an application well, we need to cache content and deliver it as a response. If the cache is not valid, it is our responsibility to bust it. REST services should be properly cached for scaling.

- **Scripts on demand:** Have you ever designed a REST service which serves the JavaScript files and you execute them on the fly? This code on demand is also the main characteristic REST can provide. It is more common to request scripts and data from the server.

- **Multiple layered system:** The REST API can be served from multiple servers. One server can request the other, and so forth. So when a request comes from the client, request and response can be passed between many servers to finally supply a response back to the client. This easily implementable multi-layered system is always a good strategy for keeping the web application loosely coupled.

- **Representation of resources:** The REST API provides the uniform interface to talk to. It uses a Uniform Resource Identifier (URI) to map the resources (data). It also has the advantage of requesting a specific data format as the response. The Internet Media Type (MIME type) can tell the server that the requested resource is of that particular type.

- Implementational freedom: REST is just a mechanism to define your web services. It is an architectural style that can be implemented in multiple ways. Because of this flexibility, you can create REST services in the way you wish to. Until it follows the principles of REST, your server has the freedom to choose the platform or technology.

## REST verb

There are quite a few REST verbs available, but six of them are used frequently.
They are as follows:

- GET
- POST
- PUT
- PATCH
- DELETE
- OPTIONS

| REST Verb | Action                                               | Success  | Failure  |
| --------- | ---------------------------------------------------- | -------- | -------- |
| GET       | Fetches a record or set of resources from the server | 200      | 404      |
| POST      | Creates a new set of resources or a resource         | 201      | 404, 409 |
| PUT       | Updates or replaces the given record                 | 200, 204 | 404      |
| PATCH     | Modifies the given record                            | 200, 204 | 404      |
| DELETE    | Deletes the given resource                           | 200      | 404      |
| OPTIONS   | Fetches all available REST operations                | 200      | -        |



