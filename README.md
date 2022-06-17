# Microservices Archetypes

This project documents the different types of microservices archetypes and the difference scenarios in which they're most commonly applicable. Multiple archetypes may apply simultaneously to a particular system.

### Quick intro

This guide does not affiliate itself to any particular technology or microservices definition. Although the concept what comprises a "microservice" is more or less intuitively understood, it's definition and implementation it's still in the realm of discussion.

Software engineers have been creating service isolation and independent scalability for decades without container orchestration. The advent and accessibility of container orchestrators, such as Service Fabric, Docker Swarm, Kubernetes, and others, made as such that this technique can be implemented to everyone, but it's implementation is not mandatory (or, is it? 🤔)

Anyways, let's talk about archetypes.

## Archetypes

### Service-specific scalability

When a particular service has a inherit capacity demand, either by it's resource requirements in nature (such as, operations have high memory or cpu demands) or by volume (larger number of requests), it is fairly common t

That doesn't mean

An example would be two distinct services:
- 

asdf

![Load oriented Microservices][1]

### General Scalability

One may say, for instance, that a Contact service typically will 
- Order Service - 
- Contact Service - Where customers will 

It is fair to say that the Order Service will, most of the time, receive. However, if experience service disruption in the Order Service or other component, Contact Service may be flooded with requests.

![Elastic][2]

### Dependency 
fasdfasdf


### Language

Either by design or , companies might use multiple language

This might be a case of service dependency

fasdfasdf

### Domain

This approach commonly is implicit to , but it is not always the case. You may which to (depending on what definition of "domain" is being used)

For example, if there , sub-domain

This approach can inadvertantly lead to [nanoservices](###nanoservices)

![Elastic][3]

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