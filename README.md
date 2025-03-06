# uhdy
It is for how to use AI assistant, not for implementing the service itself.</br>
Anyway it's an implementation of a location-based service.

## Development Environment
- VS Code
  - continue
- GPT-4o
- MacOS

## Tech Stacks
- Go
  - Fiber: Web Server
  - Huma: RESTful and Documentation
- PostgreSQL
- Grafana Loki & Promtail
- Prometheus & Grafana
- Docker Compose

## Why Golang is selected
Actually I've never used Go before and that's why Go is selected.</br>
I'd like to figure out how much AI helps me to develop the features that I want to implement
even though I don't know the language.</br>
Loki and Promtail are those that I've never used as well and they are recommended by AI for
gathering the logs.

## How much AI does
- It's like a junior software engineer who knows so much, but little bit out-of-date.
  - Guide it specifically, especially about how to design.
  - Check the latest version of the packages it provides
- I have to study the tools - Loki, Promtail - to fix configuration and make them work.
- The knowledge on Go is mendatory to review the source code it generated.
  - The scope I have to study is limited and guided by the generated code. So time is saved.
  - It enables me to implement the features while studying the language and packages step by step.

## TODO
- Add more services
- Add Gateway with Kong
