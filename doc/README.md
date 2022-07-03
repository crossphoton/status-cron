# **status-cron**

Service to build a Status Page

Run a cron with this binary for every minute to build data points


## **Services supported**
- HTTP
- Redis
- SQL
  - Postgres
  - MySQL
  - Oracle
- MongoDB

## Configuration

### Config THIS Job
Services can be fetched from the following modes and can be specified using the env `DB_TYPE`
- Postgres DB: `postgres`
  - ENV REQUIRED: `POSTGRES_URI=postgres://localhost:5432/postgres`
  - See [sample.sql](./sample.sql) to create tables
- JSON File: `json`
  - ENV REQUIRED: `JSON_PATH="/home/ec2-user/status-cron/config.json"`
  - See [sample.json](./sample.json) for configuration

> See [sample.env](./sample.env)

## Service Configuration

Base Config (Common among all):
```
{
    id: int (should be unique),
    name: string,
    url: string (taken as base for connection),
    type: [http | redis | sql | mongo],
    cron: string (determines whether to run or not),
    data: Object (service specific data)
}
```

### HTTP Services

```
data: {
    method: string (Request method),
    headers: map[string]string (Headers to be added),
    status: int (status to expect)
}
```

> Response checks are only done for status

### Redis Service

```
data: {
    password: string (taken from url if empty),
    timeout: int (Timeout for connection),
}
```

### SQL Service

```
data: {
    driver: string (driver type [postgres | mysql | oracle]),
    timeout: int (Timeout for connection),
}
```

### Mongo Service

```
data: {
    timeout: int (Timeout for connection),
}
```

## Response format
```
service_id: int
success: bool
reason: string (reason if failed)
cron_time: time (result cron time)
```

## MISC

### Debug Options
```
PRINT_RESULT=<bool> (Print the results of the services requests to console)
```