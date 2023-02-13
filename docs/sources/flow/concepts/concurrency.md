---
aliases:
- ../../concepts/concurrency/
title: Concurrency
weight: 100
---

# Concurrency

```mermaid
graph LR
A[Main] --> B[Tracer]
A --> C[Flow Controller]
A --> D[HTTP Server]
A --> E[Reporter]
C --> F[Component 1]
C --> G[Component 2]
C --> ...
```