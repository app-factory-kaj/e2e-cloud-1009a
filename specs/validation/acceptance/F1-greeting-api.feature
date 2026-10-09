Feature: F1 Greeting API

  @story-F1.1
  Rule: A named caller receives a greeting addressed to them

    Scenario: Greeting a given name
      Given the Greeting API is running
      When an API Client sends a GET request to /hello with name "Ada"
      Then the response is a JSON greeting addressed to "Ada"

  @story-F1.2
  Rule: A caller who gives no name still receives a greeting, addressed to "World"

    Scenario: Greeting with no name given
      Given the Greeting API is running
      When an API Client sends a GET request to /hello with no name parameter
      Then the response is a JSON greeting addressed to "World"

    Scenario: Greeting with an empty name given
      Given the Greeting API is running
      When an API Client sends a GET request to /hello with name ""
      Then the response is a JSON greeting addressed to "World"
