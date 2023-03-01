### Golang gRPC User microservice example with Prometheus, Grafana monitoring and Jaeger opentracing

#### Tech stack:
* [GRPC](https://grpc.io/) - gRPC
* [sqlx](https://github.com/jmoiron/sqlx) - Extensions to database/sql.
* [pgx](https://github.com/jackc/pgx) - PostgreSQL driver and toolkit for Go
* [go-redis](https://github.com/go-redis/redis) - Redis client for Golang
* [zap](https://github.com/uber-go/zap) - Logger
* [validator](https://github.com/go-playground/validator) - Go Struct and Field validation
* [migrate](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library.
* [testify](https://github.com/stretchr/testify) - Testing toolkit
* [gomock](https://github.com/golang/mock) - Mocking framework
* [Docker](https://www.docker.com/) - Docker
* [Prometheus](https://prometheus.io/) - Prometheus
* [Grafana](https://grafana.com/) - Grafana
* [Jaeger](https://www.jaegertracing.io/) - Jaeger tracing

### Usage local:
```bash
make local_up   // run containers without application
```
```bash
make run        // run the application
```
### Stop containers:
```bash
make local_down      // stop containers
```

### Usage docker:
```bash
make docker_up   // run all containers with application
```

### Stop containers:
```bash
make docker_down      // stop containers
```

### Jaeger UI:

http://localhost:16686

### Prometheus UI:

http://localhost:9090

### Grafana UI:

http://localhost:3000