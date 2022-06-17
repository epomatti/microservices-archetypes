# Microservices Archetypes

This project documents the different types of microservices archetypes and the difference scenarios in which they're most commonly applicable. Multiple archetypes may apply simultaneously to a particular system.

## Archetypes

### Service-specific scalability

When a particular service has a inherit capacity demand, either by it's resource requirements in nature (operations are heavy) or by volume (number of requests), it is fairly common t

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

### Domain

This approach commonly is implicit to , but it is not always the case. You may which to (depending on what definition of "domain" is being used)

For example, if there , sub-domain

### Team Segregation

This requirement might manifest from purely team organization, a different department or cost center, or even a 3rd party.

![Elatic][4]

### 

### Nanoservices

### "One App" Microservice




[1]: assets/load.png
[2]: assets/elastic.png
[4]: assets/team.png