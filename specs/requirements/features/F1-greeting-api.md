# Greeting API

## Purpose

Lets any API Client fetch a JSON greeting for a given name over a single,
unauthenticated HTTP endpoint, with no data persisted between requests.

## User Stories

- F1.1 As an API Client, I send a GET request to /hello with a `name` query parameter and receive a JSON greeting addressed to that name.
- F1.2 As an API Client, I send a GET request to /hello without a `name` (or an empty one) and still receive a successful JSON greeting, addressed to "World".

## Decisions

- A missing or empty `name` defaults to "World" rather than returning an error.
- The response body is `{"message": "Hello, <name>!"}`.
- `name` is echoed exactly as given, with no length limit and no sanitization beyond safe JSON encoding.