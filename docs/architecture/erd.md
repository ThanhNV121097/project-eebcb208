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

## Story extension — Display database greeting
Reviewed UI mock module `code/frontend/lib/mock/display-database-greeting.ts` exports:

```ts
export type GreetingResponse = {
  text: string;
};

export const greetingResponse: GreetingResponse = {
  text: 'Hello Word',
};
```

Existing `greetings.text` fully supplies API field `text`. No new table, column, foreign key, or index needed for this story.

## Constraints and indexes
- `greetings.id` primary key serves lookup query `SELECT text FROM greetings WHERE id = 1`.
- `greetings.id = 1` enforced by check constraint to keep one logical current greeting row.
- `greetings.text` check constraint enforces length 1..500.
- No additional index needed: primary-key lookup covers only read query.

## Migration plan

Forward:
1. Create `schema_migrations` table if absent.
2. Create `greetings` table if absent with columns and constraints listed above.
3. Insert seed row `(id = 1, text = 'Hello Word')` only when row `id = 1` does not exist.

Backward:
1. Drop `greetings` table.
2. Drop matching migration record from `schema_migrations`.

Safety on populated tables:
- Forward migration is safe on populated database: `CREATE TABLE IF NOT EXISTS` is no-op, seed insert preserves existing row and operator-updated text.
- Backward migration is destructive because it drops stored greeting data; use only for full rollback before data must be preserved.
