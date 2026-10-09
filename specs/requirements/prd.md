# Greeter

## Problem Statement

Teams building and testing services on this platform need a minimal, dependable
HTTP endpoint to greet a caller by name — something small enough to stand up
instantly as a reference or smoke-test target, with nothing else to configure
or depend on.

## Solution

Greeter is a small Go HTTP service exposing a single endpoint that returns a
JSON greeting for a given name. It follows the platform's reference service
conventions, holds no data beyond a single request, and requires no sign-in.

## Actors

- **API Client** — any caller (human, script, or another service) that sends
requests to the greeting endpoint. No identity or role is distinguished.

## Features

- F1 [Greeting API](features/F1-greeting-api.md)

## Product-wide

See [Product-wide](product-wide.md).

## Out of Scope

- Authentication or sign-in of any kind (no Thunder integration). \[org default override — explicitly excluded by the brief\]
- Persistent storage of any kind; the service keeps no data between requests.
- Additional endpoints or capabilities beyond the single greeting lookup.

## Open Questions

None — the brief fully specifies this product's scope.