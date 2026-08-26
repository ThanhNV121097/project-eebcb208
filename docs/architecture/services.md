# Service Design — greeting

## Base rules
- Routes are mounted without `/api`; public paths start with `/v1`.
- Responses are JSON unless endpoint says otherwise.
- Shared error envelope:

```json
{
  "error": {
    "code": "internal_error",
    "message": "Something went wrong."
  }
}
```

## Endpoints

### `GET /healthz`
Checks process, migrations, and database connectivity.

Response `200 text/plain`:
```text
ok
```

Response `503 application/json`:
```json
{
  "error": {
    "code": "unhealthy",
    "message": "Service is not healthy."
  }
}
```

### `GET /v1/greeting`
Returns current stored greeting.

Request body: none.

Response `200 application/json`:
```json
{
  "text": "Hello Word"
}
```

Errors:
| Status | Code | When |
|---|---|---|
| 404 | `greeting_not_found` | Required greeting row is missing |
| 500 | `internal_error` | Database or unknown server failure |
