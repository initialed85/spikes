# kcp-go_testing

## Context

Trying to build an alternative to DDS (pub/sub) that uses free software and well adopted standards where possible.

## Proposition

Use kcp-go as the messaging transport.

## Pros

- Secure
- Fairly simple
- Seems to have some native other language options

## Cons

- Getting some weird corruption stuff when I spam messages (kind of like TCP message concatenation)
    - Probably manageable
