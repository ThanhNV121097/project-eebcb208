# Design System — hello-word-D

> Source of truth: the approved `index.html` (preview: approved design).
> Every value below is extracted from it. Changing a value here without changing the approved design is a defect.

Last updated: 2025-02-14

## 1. Foundations

### 1.1 Color

Semantic tokens. Name by job, never by hue.

| Token | Value | Used for |
|---|---|---|
| `--color-bg` | `#FFFFFF` | Page background |
| `--color-surface` | `#FFFFFF` | Surface content area; same as page background in this design |
| `--color-surface-raised` | `#FFFFFF` | Not used; reserved by system tokens |
| `--color-border` | `#000000` | Not used as visible border in this design |
| `--color-text` | `#000000` | Body text, greeting |
| `--color-text-muted` | `#666666` | Hint text |
| `--color-primary` | `#000000` | Not used; reserved by system tokens |
| `--color-primary-text` | `#FFFFFF` | Not used; reserved by system tokens |
| `--color-success` | `#000000` | Not used |
| `--color-warning` | `#000000` | Not used |
| `--color-danger` | `#000000` | Not used |
| `--color-focus` | `#000000` | Focus ring |

#### Contrast audit

| Foreground | Background | Ratio | Passes |
|---|---|---|---|
| `--color-text` | `--color-bg` | `21:1` | AA / AA Large |
| `--color-text-muted` | `--color-bg` | `5.74:1` | AA |
| `--color-focus` | `--color-bg` | `21:1` | AA |

### 1.2 Spacing

Base unit: `8px`. Product uses `16px` and `32px`.

| Token | Value |
|---|---|
| `--space-2` | `16px` |
| `--space-4` | `32px` |

### 1.3 Typography

Font families:

- Body: `Arial, Helvetica, sans-serif`
- Headings: `Arial, Helvetica, sans-serif`
- Mono: not used

| Token | Size | Line height | Weight | Used for |
|---|---|---|---|---|
| `--text-3xl` | `clamp(3rem, 12vw, 8rem)` | `0.95` | `700` | Main greeting heading |
| `--text-sm` | `12px` | `1` | `400` | Hint text |

Heading levels are used in order and never skipped for visual sizing. Only `h1` appears.

### 1.4 Radius, border, shadow, motion

| Token | Value | Used for |
|---|---|---|
| `--radius-sm` | `0` | Not used |
| `--radius-md` | `0` | Not used |
| `--radius-lg` | `0` | Not used |
| `--radius-full` | `0` | Not used |
| `--border-width` | `0` | No visible borders |
| `--shadow-sm` | `none` | Resting surface |
| `--shadow-md` | `none` | Not used |
| `--shadow-lg` | `none` | Not used |
| `--duration-fast` | `0s` | No animated transitions |
| `--duration-base` | `0s` | No animated transitions |
| `--easing` | `linear` | No animated transitions |

Motion respects `prefers-reduced-motion: reduce`: state changes remain, movement is removed. This design has no motion.

### 1.5 Layout and breakpoints

| Name | Min width | Container | Columns | Gutter |
|---|---|---|---|---|
| `sm` | `0px` | full width | `1` | `32px` |
| `md` | `768px` | full width | `1` | `32px` |
| `lg` | `1024px` | full width | `1` | `32px` |
| `xl` | `1280px` | full width | `1` | `32px` |

Z-index scale (only these values are allowed):

| Layer | Value |
|---|---|
| Base | `0` |
| Sticky header | `0` |
| Dropdown | `0` |
| Modal backdrop | `0` |
| Modal | `0` |
| Toast | `0` |

## 2. Components

One subsection per reusable component. Every component lists all states.

### 2.1 Greeting Screen

**Purpose** — Full-screen landing view for exact greeting and hint.

**Anatomy** — `[main] [h1 greeting] [fixed hint]`

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default | `--color-bg`, `--color-text`, `--color-text-muted` | Single-screen greeting |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Centered black heading on white background, muted hint in bottom-left | `--color-bg`, `--color-text`, `--color-text-muted` |
| Hover | No change | None |
| Focus (keyboard) | Visible 3px black focus ring with 4px offset | `--color-focus` |
| Active / pressed | No interactive affordance | None |
| Disabled | Not applicable | None |
| Loading | Not used | None |
| Error | Not used | None |
| Empty | Not used | None |

**Accessibility** — `main` landmark, `h1` for page title, hint `aria-hidden`, keyboard focus visible on any focused element, minimum hit target not applicable.

## 3. Content and formatting

- Voice and tone: minimal, direct, neutral.
- Date, time, number, and currency formats: not used.
- Capitalization rule for buttons, headings, and labels: sentence case; greeting keeps exact text `Hello Word`.
- Empty-state and error-message wording pattern: not used.

## 4. Known deviations

| Where | Deviation | Why it stands | Follow-up |
|---|---|---|---|
| Layout | No responsive container system, no cards, no buttons, no forms | Approved design only has one centered greeting screen | None |
| Foundations | Most system tokens are unused | Smallest useful system for single-screen product | None |
| Z-index | All layers at `0` | No overlapping UI in approved design | None |

## 5. Change log

| Date | Change | Design PR |
|---|---|---|
| 2025-02-14 | Initial design system from approved greeting screen | pending |
