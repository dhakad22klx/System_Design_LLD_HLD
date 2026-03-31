# Core Software Design Principles

A quick reference guide to four foundational principles every developer should know.

---

## 1. DRY — Don't Repeat Yourself

> "Every piece of knowledge must have a single, unambiguous, authoritative representation within a system."
> — *The Pragmatic Programmer*

Avoid duplicating logic. If you find yourself writing the same code in two places, extract it.

### Violates DRY

```go
func calculateCircleArea(r float64) float64 {
    return 3.14159 * r * r
}

func calculateCirclePerimeter(r float64) float64 {
    return 2 * 3.14159 * r // Pi duplicated!
}
```

### Follows DRY

```go
const Pi = 3.14159

func calculateCircleArea(r float64) float64 {
    return Pi * r * r
}

func calculateCirclePerimeter(r float64) float64 {
    return 2 * Pi * r
}
```

---

## 2. KISS — Keep It Simple, Stupid

Good code should be:

- **Easy to read** — Another developer can understand what the code does without spending 30 minutes tracing through abstractions.
- **Easy to understand** — The logic flows naturally. No surprises, no hidden side effects, no clever tricks.
- **Easy to change** — When requirements shift, you can modify the code confidently without worrying about breaking something three layers deep.

### Violates KISS

```go
func isEven(n int) bool {
    return n&1 == 0 // Bitwise trick — clever but unclear
}
```

### Follows KISS

```go
func isEven(n int) bool {
    return n%2 == 0 // Simple and immediately obvious
}
```

---

## 3. YAGNI — You Aren't Gonna Need It

> "Always implement things when you actually need them, never when you just foresee that you need them."

Don't build features or abstractions for hypothetical future requirements. Write only what is needed *right now*.

### Violates YAGNI

```go
// We only need to greet users, but we're adding
// language support "just in case" it's needed later.
type Greeter struct {
    language string
    timezone string
    format   string
}

func (g *Greeter) Greet(name string) string {
    // Complex logic no one asked for...
    return "Hello, " + name
}
```

### Follows YAGNI

```go
func greet(name string) string {
    return "Hello, " + name
}
```

Add complexity only when a real requirement demands it.

---

## 4. Law of Demeter (LoD) — Principle of Least Knowledge

> "Only talk to your immediate friends."
> — *Law of Demeter*

A function should only call methods on:
- Itself
- Its direct parameters
- Objects it creates
- Its own fields

Avoid chaining through objects you don't directly own.

### Violates LoD

```go
// Reaching deep into an object chain
discount := order.GetCustomer().GetProfile().GetDiscount()
```

### Follows LoD

```go
// Order exposes only what callers need
discount := order.GetDiscount()

// Internally, Order handles its own traversal
func (o *Order) GetDiscount() float64 {
    return o.customer.profile.discount
}
```

---

## Summary

| Principle | Core Idea |
|-----------|-----------|
| **DRY** | Single source of truth — no duplicated logic |
| **KISS** | Simplicity over cleverness |
| **YAGNI** | Build what you need now, not what you might need later |
| **LoD** | Only interact with your direct dependencies |

> These principles work best together. DRY reduces redundancy, KISS keeps code readable, YAGNI prevents over-engineering, and LoD limits tight coupling.