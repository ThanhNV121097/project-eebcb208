# SRS — greeting

Module: `greeting`
Last updated: 2025-08-15
Design: [View the approved design](http://localhost:8080/design/eebcb208-9bfb-4de0-958d-edbd68877b81)
Design system: `design/design-system.md`

## 1. Purpose

This module shows one greeting screen for guest users. It keeps visible text driven by stored data so the page can change when the greeting row changes, without changing the screen layout. Without this module, the product has no public home page and no way to show the stored greeting.

## 2. Actors

| Actor | Who they are | What they may do in this module |
|---|---|---|
| Guest | Any visitor with no sign-in | View greeting screen and read current greeting |
| Operator | Anyone who can update stored greeting data outside this module | Change stored greeting value so the page shows the new text after reload |

## 3. Scope

**In scope** — the functions specified below, by their plan titles:

- Display database greeting

**Out of scope** — name what a reader would reasonably expect here and say where it lives instead.

- Greeting editing UI — deliberately not built; no user-facing edit flow exists in this project.
- Styling variants, animation, or alternate screens — belongs to design, but not in approved design for this module.

## 4. Functional requirements

### 4.1 Display database greeting

**Requirement GREETING-001 — Show stored greeting on page load**

*As a* Guest, *I want to* see the stored greeting on the public page, *so that* the page reflects current database content.

Behaviour:

1. When the Guest opens the greeting page, the page shows one centered line of black text on white background.
2. The text value comes from the current stored greeting value.
3. The seed greeting value is exactly `Hello Word`.
4. When the stored greeting value changes and the Guest reloads the page, the visible text matches the updated stored value.

**Acceptance criteria** — each maps one-to-one onto a test case in `docs/greeting/test-cases/display-database-greeting.md`.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | stored greeting is `Hello Word` | Guest opens page | centered greeting text `Hello Word` is visible on white background |
| AC-2 | stored greeting is changed to another value | Guest reloads page | visible text matches changed stored value |
| AC-3 | greeting page is open | Guest inspects layout | exactly one centered greeting line is shown, with black text on white background |

**Failure, boundary and permission behaviour**

| Case | Condition | Expected behaviour |
|---|---|---|
| Invalid input | Not applicable: no user input exists in approved design for this screen |
| Boundary | Not applicable: design shows one fixed-screen layout with one greeting line |
| Not found | Not applicable: approved design has no empty-state screen |
| Not permitted | Not applicable: no permissions are shown or required for guest read |
| Conflict | Not applicable: one reader screen with no concurrent edit UI |
| Upstream failure | Not applicable in approved design: no error state is drawn; API error handling belongs to service contract and backend implementation |

**Data touched** — the fields this function reads and writes, in product terms.

| Field | Type | Required | Rule |
|---|---|---|---|
| Greeting text | text | yes | Must exist and render as exact visible text; seed value is `Hello Word` |

## 5. Screens

| Screen | Section in the design | Functions it serves | States that must exist |
|---|---|---|---|
| Greeting screen | Main centered greeting view | GREETING-001 | default |

## 6. Non-functional requirements

| Area | Requirement |
|---|---|
| Accessibility | Greeting text is readable with contrast ≥ 4.5:1 against white background, and page focus ring remains visible on keyboard focus |
| Responsive | Screen keeps centered layout at 320px viewport width and up, with no horizontal page scroll |
| Performance | Initial page render shows greeting within 2s at p95 on a 1 Mbps connection with cold cache |

## 7. Dependencies and assumptions

- **Depends on:** PostgreSQL, for storing the greeting value.
- **Depends on:** backend HTTP API, for reading the current greeting value.
- **Assumption:** Only one greeting value exists at a time; if multiple rows are introduced, this module still renders one current greeting and TL must define selection rules.

| Open question | Proposed default | Who decides |
|---|---|---|
| None for current scope | N/A | N/A |

## 8. Traceability

| Plan item | Requirement ids | Test cases |
|---|---|---|
| Display database greeting | GREETING-001 | `test-cases/display-database-greeting.md` |
