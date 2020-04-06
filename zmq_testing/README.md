# zmq_testing

## Context

Trying to build an alternative to DDS (pub/sub) that uses free software and well adopted standards where possible.

## Proposition

Use ZeroMQ as the messaging transport.

## Pros

- Nice abstractions that hide away the complexity of socket communication
- Performant

## Cons

- Most implementations require underlying C lib (not Pure Go / Pure other languages)
- No support for UDP (except for PGM, I assume)
