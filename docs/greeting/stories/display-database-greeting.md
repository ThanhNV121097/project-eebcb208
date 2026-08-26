# Story — Display database greeting

## User story
As a Guest, I want to see current stored greeting on public page, so that page reflects database content.

## In scope
- One public greeting screen for Guest.
- Read current greeting from backend data source on page load.
- Show exact stored text centered on white background.
- Show seed text `Hello Word` on first load when seed row is unchanged.
- Refreshing page after greeting value changes shows updated stored text.

## Out of scope
- Greeting editing UI.
- Sign-in, permissions, or any operator tools.
- Multiple greeting selection rules.
- Alternate screens, visual variants, or animation.
- Error or empty-state UI beyond approved single-screen design.

## UI scope
- Greeting screen only, matching approved design.
- Default state only: one centered greeting line, black text on white background.
- No extra controls, no additional layout states, no loading or error view in UI.

## Acceptance criteria
1. Given stored greeting is `Hello Word`, when Guest opens page, then centered greeting text `Hello Word` is visible on white background.
2. Given stored greeting changes to another value, when Guest reloads page, then visible text matches changed stored value.
3. Given greeting page is open, when Guest inspects layout, then exactly one centered greeting line is shown with black text on white background.
4. Given page loads normally, when greeting is rendered, then no edit controls, sign-in prompts, or alternate screens appear.

## Dependencies
- PostgreSQL available with greeting row seeded to `Hello Word`.
- Backend HTTP API available to read current greeting value.
- Approved design and design system remain unchanged.

## Notes
- This story covers display only; data update behavior lives outside this module.
