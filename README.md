# Mathematical Groups in Go

An implementation of mathematical groups in Go, designed by a geometric group theorist.

## Motivation

I’m an Oxford mathematics graduate who specialized in geometric group theory. This project has several aims, the most personal of which is to build a bridge between my mathematical background and my programming skills.

There’s a good reason for that: I loved group theory, but programming is, at present, far more directly useful to society. I hope this project can serve as a bridge for other people too, and encourage more practical applications of the beautiful theory of groups. Groups are extremely common objects in mathematics, so there’s plenty of untapped potential!

With the arrival of generics in Go 1.18 (2022), it finally became realistic to model mathematical structures in a way that feels natural. This project is my attempt to bring the world of group theory into a modern, strongly typed language, in a style that could plausibly fit a "math/group" package in the Go standard library.

## Quickstart

This project is currently a **library**, not a standalone tool. To use it, import it into your Go code.

A CLI tool is planned for the future; once it exists, I’ll add installation and usage instructions here.

**Important:** verifying group axioms (associativity, identity, inverses) and subgroup properties (e.g. normality), etc., is left to the user. The focus of this project is on performing group-theoretic computations that would be tedious by hand, not on theorem or property verification (which is better suited to systems like Lean).

## Dependencies

For the core libraries, I’m not using any external dependencies. From the Go standard libary, I use only the errors package so far.

The CLI will use the standard library, of course. Once the CLI incorporates SQL (Postgres + migration / query tooling), I’ll list those dependencies here.

## Roadmap

This roadmap is not fixed and not necessarily in implementation order. Items may be split further as the project grows.

- [ ] Free groups
- [ ] Group presentations
- [ ] CLI tool
- [ ] Finite groups
- [ ] Word problem solvers (for certain classes of groups)
- [ ] Subgroups
- [ ] Homomorphisms
- [ ] Automatic groups
- [ ] Varieties of groups
- [ ] SQL-backed storage for group data
- [ ] And more as the project evolves

## Contributions

This project is currently a personal exploration, and I’m **not accepting external pull requests** for now.

However, if you spot a bug, please open an issue on GitHub.

Suggestions that stay strictly within group theory are most welcome. For example, a request to add matrix groups over a given ring may fit, but I don’t plan to implement substantial ring theory here (that would belong in a separate project!).

