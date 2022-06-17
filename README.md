# Microservices Archetypes

This project documents the different types of microservices archetypes and the difference scenarios in which they're most commonly applicable. Multiple archetypes may apply simultaneously to a particular system.

### Quick intro

This guide does not affiliate itself to any particular technology or microservices definition. Although the concept what comprises a "microservice" is more or less intuitively understood, it's definition and implementation it's still in the realm of discussion.

Software engineers have been creating service isolation and independent scalability for decades without container orchestration. The advent and accessibility of container orchestrators, such as Service Fabric, Docker Swarm, Kubernetes, and others, made as such that this technique can be implemented to everyone, but it's implementation is not mandatory (or, is it? 🤔)

## Archetypes

### Service-specific Scalability

When a particular service has unique capacity or operational requirements, engineers might choose to implement service isolation. Such decision can be driven by multiple reasons:

- **Business criticality** - Business operation or reputation is highly sensitive to the operation of such service. This could be a main service that a direct dependency for the purchase of services and goods by customers, or even compliance related, where govern or institutions may fine the company if the operation doesn't meet contract targets.
- **Performance** - When resource consumption requirements re high nature in (operations require high memory, cpu, throughput, etc).
- **Volume** - The service receives a distinctively high amount of transactions or tasks to be completed, either constantly or during peak events.

![Load oriented Microservices][1]

### General Scalability

Typically all components are or should be designed to scale in a system, with specific scalability configuration being applied to individual services.

This is a constant activity guided by observation and testing via APM techniques, and maintained as the system evolves across development cycles.

![Elastic][2]

Large systems will scale this model to the extreme across multiple products and clusters to reach customers globally, often arriving in unique architectures with multiple layers of dependencies and abstractions, although that would not classify as "micro" services architecture by most definitions.


### Technology Requirement 

When there is a technical requirement that forces a service to be created using a specific technology. Reasons may vary as to why this would be the case:

- An SDK required for integration with a 3rd party is available in only a few languages
- A software dependency that needs to run on a specific operating system
- A performance requirement that may be achieved only or more easily by a specific technology

You may have you LoB services with a standard template, and eventually require some of them to be built in different technologies.

![Elastic][6]


### Language

Either by design or , companies might use multiple language

This might be a case of service dependency

You _can_ use use multiple languages for difference microservices, but that doesn't necessarily means that you _should_ do it.



### Technical Microservices

asdfsadf
### Domain

This approach commonly is implicit to , but it is not always the case. You may which to (depending on what definition of "domain" is being used)

For example, if there , sub-domain


![Elastic][3]

This approach can inadvertantly lead to [nanoservices](###nanoservices).

### Team Segregation

This requirement might manifest from purely team organization, a different department or cost center, or even a 3rd party.

![Elatic][4]

### Migration

Gradual migration from a legacy application to a microservices architecture might be implemented, where new services are created in isolation, while existing services can be gradually migrated to the platform.

![Elatic][5]

### Nanoservices

In smaller projects where decision to implement microservices, either in

### "One App" Microservice

This is simply an single application that is 


[1]: assets/load.png
[2]: assets/elastic.png
[3]: assets/domain.png
[4]: assets/team.png
[5]: assets/migration.png
[6]: assets/tech.png