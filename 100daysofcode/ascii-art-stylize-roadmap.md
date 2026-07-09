# ASCII-Art-Stylize: 14-Day Learning Roadmap

**Project:** ascii-art-stylize (Go + CSS)
**Budget:** 4 hours/day x 14 days = 56 hours total
**Prior CSS experience:** None

---

## Core Insight From Planning

I was initially overestimating how much CSS I needed because I was counting spec **requirements** instead of underlying **mechanisms**. The same mechanism (e.g. `:hover`, `:focus`) satisfies multiple requirements at once. So the plan should revolve around a small set of core, reusable concepts rather than a long list of isolated topics.

The bigger lesson: my real question shifted from

> "What CSS should I learn in 56 hours?"

to

> "What mechanisms does this specification require, and how can I verify I truly understand each one while building the real project?"

This is a transferable method for approaching **any** new spec or technology:

1. Read the spec for underlying mechanisms (not just requirements as a list)
2. Form a hypothesis about what's load-bearing
3. Test the hypothesis against the real project
4. Only reach for a new tool when the problem actually demands it

---

## Spec → Mechanism Mapping (my own analysis)

**"More appealing, interactive, intuitive"**
→ typography, spacing/layout, hover/focus states, cursor, transitions, visual hierarchy

**"Give more feedback to the user"**
→ hover / active / focus / disabled states, loading indicators, success/error messaging, input validation styling

**"Text properly visible regardless of colors used"**
→ NOT just "pick nice colors" — accessibility / contrast ratios (WCAG standard):
- 4.5:1 minimum contrast for normal-sized text
- 3:1 minimum contrast for large text (~18pt regular / 14pt bold)

Also affected by: font size, weight, line spacing, background choice.

> **Open question:** does this requirement apply only to the UI chrome, or also to the rendered ASCII art output itself (if per-character color is applied)? Resolve by checking my own `ascii-art-color` code: does `--color` wrap individual matched substrings, or color the whole block? That answer likely settles the scope on its own.

**"Consistent, responsive, interactive"**
- *Consistent* = a discipline, not a new topic — reuse the same spacing, colors, fonts everywhere. Best enforced via **CSS custom properties (variables)** defined once (e.g. in `:root`) and reused everywhere — turns consistency into something structural rather than something enforced by eye/memory.
- *Responsive* = adapting layout to screen width. My project's HTML is basically a single-column form + output block, so likely does **not** need full Grid/complex Flexbox — minimum viable approach: centered container, max-width, vertical flow, maybe 1-2 media queries, horizontal scroll for the `<pre>` output block if needed.
- *Interactive* = same pseudo-class states as "feedback" above.

**Load-bearing concept:** the CSS box model + normal document flow (margin, padding, width, display). Almost everything else — typography, responsive layout, interactive states, even accessibility via spacing — depends on understanding this first. Hypothesis to test on Day 1–2.

---

## User Journey / Feedback Gap Audit (ascii-art-web, current state)

| Moment | Current feedback | Gap |
|---|---|---|
| Landing on page | None | No visual hierarchy |
| Typing text | None | No focus border / highlight / placeholder styling |
| Choosing options | Browser default only | No hover/selected styling |
| Pressing Submit | None | No disabled/loading state |
| Waiting for response | None | No loading indicator |
| Output appears | Instant, unstyled | No transition, no `<pre>` styling |
| Error occurs | Plain/unclear | No consistent error styling |

Each gap = candidate for a pseudo-class or small HTML/CSS addition.

---

## Roadmap

### Phase 0 — Setup (Day 1, ~30 min)
- Bookmark [WebAIM Contrast Checker](https://webaim.org/resources/contrastchecker/)
- Open `ascii-art-web` unstyled HTML in browser + DevTools open throughout
- [Chrome DevTools CSS overview](https://developer.chrome.com/docs/devtools/css)

### Phase 1 — Days 1–2: Box model & document flow (experiment-driven)
**Goal:** test the "box model is load-bearing" hypothesis.

**Constraint:** only `margin`, `padding`, `width`/`max-width`, `display`, `font-family`, `font-size`, `line-height`, `text-align`, `borders`, `background-color`. No Flexbox, Grid, animations, or media queries yet.

**Question to answer:** how close (measured against the checklist below, not a vague "feeling") can I get to satisfying the spec using only fundamentals?

Resources:
- [MDN: The box model](https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/The_box_model)
- [MDN: Normal flow](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_display/Normal_Flow)
- Video: [Kevin Powell — "The CSS Box Model"](https://www.youtube.com/watch?v=rIO5326FgPE)

### Phase 2 — Days 3–4: Typography, spacing, + CSS custom properties
Define reusable design tokens early (spacing scale, 2–3 font sizes, color palette) using CSS variables, since consistency depends on this.

**Question to test:** where must a custom property be defined (e.g. `:root`) for all components to see it? What happens if scoped to one component instead of globally? Predict before testing.

Resources:
- [MDN: Using CSS custom properties](https://developer.mozilla.org/en-US/docs/Web/CSS/Using_CSS_custom_properties)
- [web.dev: Learn CSS — Typography module](https://web.dev/learn/css/typography)
- Video: [Kevin Powell — "CSS Variables"](https://www.youtube.com/watch?v=Qf1_hs4BC8g)

### Phase 3 — Day 5: Flexbox (only if still needed)
Re-run the Phase 1 checklist first. Only learn Flexbox for the **specific** layout problem that document flow + box model couldn't solve.

Resources:
- [CSS-Tricks: A Guide to Flexbox](https://css-tricks.com/snippets/css/a-guide-to-flexbox/)
- [Flexbox Froggy](https://flexboxfroggy.com/) (interactive, ~30 min max)

### Phase 4 — Day 6: Responsive design (minimal scope)
Resources:
- [web.dev: Responsive Design basics](https://web.dev/learn/design/)
- [MDN: Using media queries](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_media_queries/Using_media_queries)

Test using DevTools device toolbar, not manual window resizing.

### Phase 5 — Days 7–10: Component-by-component build
For each component (input, select, button, `<pre>` output block): style layout + interactive states + accessibility **together**, not sequentially in separate phases. This mirrors how real frontend developers actually build interfaces.

Resolve the ASCII-output-color accessibility scope question before styling the output component.

Resources:
- [MDN: Pseudo-classes (:hover, :focus, :active, :disabled)](https://developer.mozilla.org/en-US/docs/Web/CSS/Pseudo-classes)
- [MDN: :focus-visible](https://developer.mozilla.org/en-US/docs/Web/CSS/:focus-visible)
- [A11y Project checklist](https://www.a11yproject.com/checklist/) (run per component, not just at the end)

### Phase 6 — Days 11–14: Consistency audit, responsiveness re-check, review
- Re-run the Phase 2 checklist against the finished app
- Look for inconsistency: do buttons look related? Is spacing predictable? Does every interactive element behave similarly?
- Re-test responsiveness across breakpoints
- Remove unnecessary CSS, refactor
- If possible, get someone else to use it and watch for confusion

Optional: [Refactoring UI](https://www.refactoringui.com/) (paid) — skim relevant sections if time allows

---

## General / Reference Resources

- [MDN CSS docs](https://developer.mozilla.org/en-US/docs/Web/CSS) (reference, not cover-to-cover)
- [web.dev Learn CSS](https://web.dev/learn/css/) (structured path)
- [Laws of UX](https://lawsofux.com/) (short, visual, high signal)
- [freeCodeCamp Responsive Web Design curriculum](https://www.freecodecamp.org/learn/2022/responsive-web-design/)
- YouTube: **Kevin Powell** (primary video resource throughout)
- YouTube: freeCodeCamp.org — Responsive Web Design Full Course
- YouTube: Traversy Media — Flexbox / CSS Grid Crash Courses

---

## Day 1 Checklist (concrete first-session plan)

1. Open current `ascii-art-web` project
2. Open it in the browser
3. Open DevTools, keep it open throughout the session
4. Bookmark + open WebAIM Contrast Checker
5. Read MDN pages on Box Model and Normal Flow
6. Immediately apply concepts to the real project (no isolated exercises unless genuinely stuck)
7. Keep a short journal: *"I added X because..."* for every change — this is how I catch imitation vs. understanding in the moment

> **Rule:** Replace the instinct "how do I make X work" with "why is the browser doing this?"
> If stuck >15–20 min with no progress, that's a signal to look up just enough vocabulary to reformulate the question — not a failure of the rule.

---

## End-of-Session Checkpoint (ask every day, every 4-hour block)

1. Could I rebuild today's work without reopening the tutorial?
2. Can I explain why every CSS rule I added exists?
3. Can I predict what would happen if I changed one of those rules?
4. Did I remove any unnecessary CSS? (If deleting it changes nothing, I didn't understand why I added it.)
5. Did I verify accessibility instead of assuming it?

If "no" to more than one → don't rush to the next topic.

**Ultimate test:** *"If someone deleted everything I wrote today but left me the same HTML, could I rebuild it from understanding instead of memory?"*

---

## "70–80%" Experiment Scoring Checklist (Phase 1 measurement)

Score each independently — don't average into a vague feeling:

- [ ] **Appealing** — clear spacing between sections? readable typography? does the page feel organized?
- [ ] **Interactive** — is it clear what's clickable? do interactive elements have obvious states?
- [ ] **Feedback** — does every important action communicate something? can the user tell what's happening?
- [ ] **Text visibility** — is text comfortably readable? does it meet contrast requirements (check against WebAIM checker)?
- [ ] **Responsive** — usable on desktop? usable on mobile without horizontal scroll (except possibly the ASCII output)?

---

## Open Questions / To Resolve Before Phase 5

- Does "text visible regardless of colors used" apply to the rendered ASCII art output, or only to UI chrome (forms/buttons/labels)?
  - Check: my own `ascii-art-color` implementation (per-character span vs. whole-block coloring) — this alone may settle it
  - Also check: spec wording, any accepted reference projects if available

---

## Notes / Reflections

*(add your own below as you go)*

