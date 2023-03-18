## Onion architecture

### Application Structure & Layers

|Layers|Package|Depends On|Testing|Description|
|------|-------|----------|-------|-----------|
|Domain Model|model|-|unit|Responsible for domain identities and business rules|
|Domain Services|domain|model|unit|responsible for business rules and knowledge, providing interfaces for repositories|
|Application Services|service|domain|unit|responsible for orchestration, authentication, authorisation and excludes business rules, providing interfaces for external services.|
|Infrastructure Services|adapter-db adapter-stream adapter-api|model, domain, service|unit,integration|acts as an entrypoint for application services like Rest API, Soap API, Event Streams, Queues and scheduled Jobs or provides implementations for interfaces defined in domain or application services|
|Application|main|all|Integration Testing E2E Testing BDD|packaging all layers into either web application, command line application.|
|Observability services||tracing, logging, metrics etc. libs|Observability services are responsible for monitoring the application|