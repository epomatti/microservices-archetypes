# Microservices Archetypes

This project documents the different types of microservices archetypes and scenarios that I observed over the years in which they're most commonly applicable. Multiple archetypes may apply simultaneously to a particular system.

<details>
  <summary>📓 Note</summary>
  This guide does not affiliate itself to any particular technology or microservices definition, and also does not (or at very minimum tries not to) advocate XYZ techniques as good or bad.

  There's always trade offs and situations that are particular to some projects.

  Although the concept of what constitutes a "microservice" is more or less intuitively understood, it's definition and implementation still are a topic of discussion.

  Software engineers have been creating service isolation and independent scalability for decades without container orchestration. The advent and accessibility of container orchestrators, such as Service Fabric, Docker Swarm, Kubernetes, and others, allowed this technique to be available to everyone, but it's implementation is not mandatory. (Or is it? 🤔)
</details>



## Archetypes

### 1 - Service-specific Scalability

When a particular service has unique capacity or operational requirements, engineers might choose to implement service isolation. Such decision can be driven by multiple reasons:

- **Business criticality** - Business operation or reputation is highly sensitive to the operation of such service. This could be a main service that a direct dependency for the purchase of services and goods by customers, or even compliance related, where govern or institutions may fine the company if the operation doesn't meet contract targets.
- **Performance** - When resource consumption requirements re high nature in (operations require high memory, cpu, throughput, etc).
- **Volume** - The service receives a distinctively high amount of transactions or tasks to be completed, either constantly or during peak events.

![Load oriented Microservices][1]

### 2 - General Scalability

Typically all components are or should be designed to scale in a system, with specific scalability configuration being applied to individual services.

This is a constant activity guided by observation and testing via APM techniques, and maintained as the system evolves across development cycles.

![Elastic][2]

Large systems will scale this model to the extreme across multiple products and clusters to reach customers globally, often arriving in unique architectures with multiple layers of dependencies and abstractions, although that would not classify as "micro" services architecture by most definitions.


### 3 - Technology Requirement 

When there is a technical requirement that forces a service to be created using a specific technology. Reasons may vary as to why this would be the case:

- An SDK required for integration with a 3rd party is available in only a few languages
- A software dependency that needs to run on a specific operating system
- A performance requirement that may be achieved only or more easily by a specific technology

You may have you LoB services with a standard template, and eventually require some of them to be built in different technologies.

![Elastic][6]


### 4 - Multi-Language

A company or team may implement multiple languages. That might happen by design or by outer forces, such as a particular demand for professionals.

![Language][7]

Architectures _can_ use use multiple languages for difference microservices, but that will come with a cost.

Often microservices will evolve and generate complex core libraries and CI/CD templates that would require significant effort for migration. Teams should consider this approach carefully. This could inevitably lead to [technical microservices](#technical-microservices).


### 5 - Technical Microservices

Some requirements cannot be met with code dependencies and might need to be exported to their very own microservice for technical reasons.

This generate problems of it's own, but again for whatever reason a team might opt to build such a service.

Let's suppose that using [Azure App Configuration](https://azure.microsoft.com/en-us/services/app-configuration/) is not approved, and using an orchestrator config and secrets layer is not flexible enough, a team might choose to develop it's own configuration manager layer.

![Language][8]

### 6 - Domain

This approach commonly is implicit to , but it is not always the case. You may which to (depending on what definition of "domain" is being used)

For example, if there , sub-domain


![Elastic][3]

This approach can inadvertantly lead to [nanoservices](#nanoservices).

### 7 - Team Segregation

Different teams may own specific microservices.
- Cost distribution and/or internal service divisions
- Code management
- 3rd party development

This can possibly lead to a multi-language microservices architecture.

![Elatic][4]

### 8 - Migration

Gradual migration from a legacy application to a microservices architecture might be implemented, where new services are created in isolation, while existing services can be gradually migrated to the new platform.

![Elatic][5]

### 9 - Nanoservices

Often in smaller projects that implement microservices, but not exclusive to those, some or even all services are slipt into their respective domains, but the amount of code and operations are not enough to justify the inherent complexity.

During the startup period a lot of "MVPs" were developed in such a way. Is the complexity justified by the possibility of exponential growth in the future?

In one particular project that I consulted for, the architecture was almost a single function in each pod.

![Elatic][9]

### 10 - Extra: "One App" Microservice

This is simply a single application deployed to any platform or container which is mistakenly called "microservice", but has none of the scalability and boundaries that one would expect. It's just an app.

The first time a heard it was for a Java Spring application. Since they were different from Java EE in which you can port everything in a single file, that probably started the whole thing.

![Elatic][10]




[1]: assets/load.png
[2]: assets/elastic.png
[3]: assets/domain.png
[4]: assets/team.png
[5]: assets/migration.png
[6]: assets/tech.png
[7]: assets/language.png
[8]: assets/config.png
[9]: assets/nano.png
[10]: assets/app.png