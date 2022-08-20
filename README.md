### Experiment: gostack-gql

Nothing fancy. Just trying out some basic backend. The main motivation is to see if it's viable to go schema-first for a backend application.

Database functions are generated from database migration files with [sqlc](https://sqlc.dev) and the API is generated from a GraphQL schema files using [gqlgen](https://gqlgen.com).

The goal is to be able to focus on the business logic, and not re-implement things that are already defined in schemas.

### Commands

- `goose -dir pkg/pgdb/schema up`
