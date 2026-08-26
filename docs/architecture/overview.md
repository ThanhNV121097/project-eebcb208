# Architecture Overview — hello-word-D

## Scope
Fullstack greeting app: PostgreSQL stores one greeting, Go API serves it, Next.js page renders it. Seed value is exactly `Hello Word`.

## Stack
| Layer | Choice | Reason | Rejected alternative |
|---|---|---|---|
| Frontend | Next.js 15 App Router, TypeScript, Tailwind v3 | Matches repo scaffold and supports server-rendered first paint | Static HTML would not fetch stored greeting |
| Backend | Go 1.22 HTTP server | Small API, native concurrency, existing container contract | Node API would add second runtime style |
| Database | PostgreSQL 16 | Required persistence for greeting value | In-memory value would fail reload-after-change requirement |
| CI | Existing `.github/workflows/ci.yml` | Builds, vets, tests, lints, token-checks PRs | Custom workflow blocked by read-only `.github/` |

## Repository layout
```text
code/backend/                  Go service
  cmd/api/main.go              HTTP server, migration runner, health check
  internal/migrations/         embedded SQL migrations
  migrations/                  tracked SQL source for reviewers
code/frontend/                 Next.js app
  app/layout.tsx               root metadata and shell
  app/page.tsx                 composition root only
  app/globals.css              shared tokens and base styles
  components/                  story components later
docs/architecture/             TL-owned contracts
```

## Runtime flow
1. Backend reads `DATABASE_URL`, applies all embedded migrations in filename order, then starts on `PORT`, fallback `APP_PORT`, fallback `8080`.
2. `/healthz` returns 200 only after migrations have succeeded and `SELECT 1` succeeds.
3. Frontend reads greeting from backend during render. Browser-facing calls use `NEXT_PUBLIC_API_URL`; server-side render uses `API_ORIGIN`.
4. `docker compose --profile local up --build` starts PostgreSQL, backend, frontend. Existing compose file names PostgreSQL service `postgres` and volume `pg_data`.

## Naming and code conventions
- Backend routes use `/v1/...`; no `/api` prefix.
- Backend responses use shared JSON error envelope from `services.md`.
- Migrations use timestamped `YYYYMMDDHHMMSS_name.up.sql` and `.down.sql` files.
- Frontend component files use `export default function ComponentName()`.
- `app/page.tsx` stays server component and only composes story components.
- CSS values come from `app/globals.css` tokens. No CSS-module hardcoded colors, long px values, or token fallbacks.

## Environment variables
| Service | Key | Required | Purpose |
|---|---|---|---|
| backend | `DATABASE_URL` | yes | PostgreSQL connection string injected by runtime |
| backend | `PORT` | no | Primary HTTP port |
| backend | `APP_PORT` | no | Legacy fallback HTTP port |
| frontend | `NEXT_PUBLIC_API_URL` | yes | Browser-visible backend origin |
| frontend | `API_ORIGIN` | yes | Backend origin for server-side rendering |
| compose | `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB` | local only | Local PostgreSQL settings |

## Security, failure handling, observability
- External input is limited to HTTP method/path; all DB work uses parameterized queries.
- Startup fails closed when migrations or DB ping fail, so broken schema never reports healthy.
- Health endpoint returns generic status only; no secrets or DSNs in responses.
- Logs use stdout/stderr from service processes; compose and deploy capture them.

## Decisions and risks
| Decision | Why | Tradeoff |
|---|---|---|
| Self-migrate on backend boot | Runtime database starts empty | Startup depends on schema lock and DB availability |
| Embed migrations in Go binary and keep SQL files visible under `migrations/` | Container can run without extra file copy while reviewers see schema | SQL appears in two paths and must stay mirrored |
| One `greetings` row keyed by `id = 1` | SRS says only one current greeting exists | Future multi-greeting selection needs new ERD decision |
| No frontend feature component yet | Scaffold must not implement story UI | First story owns component and page mount |

## Run and check
```bash
cp .env.example .env
docker compose --profile local up --build
cd code/backend && go build ./... && go vet ./... && go test ./...
cd code/frontend && npm ci && npm run lint && npm run build && npm test --if-present
```
