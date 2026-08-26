# Test Cases — Display database greeting

Risk level: low. One read-only screen. Coverage stays on required happy path and contract shape; no failure path is specified for this module.

## Case 1 — Seed greeting renders on open
**Scenario**: AC-1 seed greeting visible
**Given** stored greeting row value is exactly `Hello Word`
**When** Guest opens greeting page
**Then** browser displays exactly one centered greeting line with visible text `Hello Word` on white background
**Check**: render_url
**Trace**: GREETING-001, AC-1

## Case 2 — Updated stored greeting shows after reload
**Scenario**: AC-2 reload picks updated stored greeting
**Given** stored greeting row was changed to `Bonjour Word`
**When** Guest reloads greeting page
**Then** browser displays visible text `Bonjour Word`
**Check**: render_url
**Trace**: GREETING-001, AC-2

## Case 3 — Layout stays one centered black line on white
**Scenario**: AC-3 centered single-line layout
**Given** greeting page is open
**When** Guest inspects page layout
**Then** browser shows exactly one greeting line, text is black, background is white, and line is centered
**Check**: measure_styles
**Trace**: GREETING-001, AC-3

## Case 4 — Greeting API returns stored text
**Scenario**: GET /v1/greeting success shape
**Given** stored greeting row value is `Hello Word`
**When** client requests `GET /v1/greeting`
**Then** response is `200 application/json` with body `{ "text": "Hello Word" }`
**Check**: fetch_url
**Trace**: service contract, GET /v1/greeting success
