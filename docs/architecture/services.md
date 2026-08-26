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

Auth: none.

Request body: none.

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

Errors:
| Status | Code | When |
|---|---|---|
| 503 | `unhealthy` | Process, migration, or database connectivity check fails |

### `GET /v1/greeting`
Returns current stored greeting.

Auth: none.

Request body: none.

Response `200 application/json`:
```json
{
  "text": "Hello Word"
}
```

Response shape matches reviewed UI mock module `GreetingResponse` exactly:

```ts
{
  text: string;
}
```

Errors:
| Status | Code | When |
|---|---|---|
| 404 | `greeting_not_found` | Required greeting row `id = 1` is missing |
| 500 | `internal_error` | Database or unknown server failure |

## Story extension — Display database greeting
No new endpoint needed beyond `GET /v1/greeting`. Backend implementation must read `greetings.text` for row `id = 1` and return only:

```json
{
  "text": "<stored greeting text>"
}
```

Frontend mock imports `greetingResponse.text`; API response keeps same field name, type, required status, and no list envelope.

## Migration plan
Service layer depends on ERD migration creating and seeding `greetings`.

Forward:
1. Apply database migration before serving traffic.
2. Start service after migration succeeds.
3. `GET /healthz` returns `200 ok` only when migrations and database connectivity pass.

Backward:
1. Stop service version that depends on `greetings`.
2. Roll back database migration as described in ERD.
3. Deploy prior service version or leave service stopped.

Safety on populated tables:
- Forward is safe because seed insert does not overwrite existing `greetings.id = 1` row.
- Backward removes stored greeting data with dropped table; treat as destructive rollback.
