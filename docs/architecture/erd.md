# ERD — greeting

## Tables

### `greetings`
| Column | Type | Constraints | Notes |
|---|---|---|---|
| `id` | integer | primary key, `id = 1` enforced | Single current greeting row |
| `text` | text | not null, length 1..500 | Visible greeting, seeded as `Hello Word` |
| `created_at` | timestamptz | not null, default `now()` | Audit timestamp |
| `updated_at` | timestamptz | not null, default `now()` | Audit timestamp |

### `schema_migrations`
| Column | Type | Constraints | Notes |
|---|---|---|---|
| `version` | text | primary key | Migration filename stem |
| `applied_at` | timestamptz | not null, default `now()` | Applied timestamp |

## Relationships
No foreign keys. `greetings` holds exactly one logical row for current greeting.

## Seed data
Migration inserts row `(id = 1, text = 'Hello Word')` and leaves later operator updates intact on re-run.
