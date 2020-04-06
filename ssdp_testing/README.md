# ssdp_testing

## Context

Trying to build an alternative to DDS (pub/sub) that uses free software and well adopted standards where possible.

## Proposition

Use SSDP as the discovery mechanism.

## Pros

- Pure Go
- An existing standard
- Seems pretty simple

## Cons

- Might have some issues on a network with lots of devices using SSDP  
- Might litter the application state of other SSDP devices
