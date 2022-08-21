### Experiment: gostack-gql

Nothing fancy. Just trying out some basic backend. The main motivation is to see if it's viable to go schema-first for a backend application.

Database functions are generated from database migration files with [sqlc](https://sqlc.dev) and the API is generated from a GraphQL schema files using [gqlgen](https://gqlgen.com).

The goal is to be able to focus on the business logic, and not re-implement things that are already defined in schemas.

## Schemas as source of truth

### Database schema: goose

1. Add a new migration file into pkg/pgdb/schema
2. Run the server (`go run server.go`)

### Data management: sqlc

1. Add new queries or commands to `pkg/pgdb/query.sql`
2. Run `go generate ./...`

### GraphQL API: gqlgen

1. Make the necessary schema changes
2. Run `go generate ./...`
