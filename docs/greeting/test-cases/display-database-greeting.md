# Test Cases — Display database greeting

Risk level: low. One read-only guest screen, no user input, no edit flow.

## Cases

**Scenario**: Show seed greeting on first load
**Given** stored greeting row exists with exact value `Hello Word`
**When** guest opens greeting page
**Then** centered visible text `Hello Word` appears on white background
**Check**: render_url
**Trace**: GREETING-001 AC-1

**Scenario**: Show updated stored greeting after reload
**Given** stored greeting row already changed to `Bonjour` before page load
**When** guest opens greeting page, then reloads page after stored value stays `Bonjour`
**Then** visible text matches current stored value `Bonjour` after reload
**Check**: render_url
**Trace**: GREETING-001 AC-2

**Scenario**: Show single centered black line on white background
**Given** greeting page is open
**When** guest inspects screen layout
**Then** exactly one greeting line is visible, text is black, background is white, and line is centered
**Check**: measure_styles
**Trace**: GREETING-001 AC-3
