# Mathematical Groups in Go

An implementation of mathematical groups in Go, designed by a geometric group theorist. This project is a personal exploration and not connected to any academic program.

## Motivation

During my master's degree in mathematics, I specialized in geometric group theory. Amazed by the deep links between theoretical computer science and geometry, I decided to build this bridge between the beautiful theory of groups and programming. With the arrival of generics in Go 1.18 (2022), it finally became realistic to model mathematical structures in a way that feels natural. This project is my attempt to bring the world of group theory into a modern, strongly typed language, in a style that could plausibly fit a "math/group" package in the Go standard library. I hope this projevt will encourage more practical applications of group theory; as groups arise naturally in many contexts, there certainly is plenty of untapped potential!

Bridging programming and group theory is highly non-trivial, and as new research discoveries are found, this project will reflect those and thus will never truly be "finished". Mathematical proofs describing an algorithm (or proving the mere existence of such an algorithm) do not establish the careful algorithmic and arhcitectural design required to implement in software. Even a single step in a mathematical proof, like word reduction or finding subwords, requires careful considerations, such as in time complexity. Other algorithms require initiating multiple processes concurrently, another reason I chose Go for this project. Therefore, one aim of this project is to translate those theoretical procedures into efficient implementations suitable for real computation.

## Quickstart

This project is currently a **library**, not a standalone tool. To use it, import it into your Go code.

A CLI tool is planned for the future; once it exists, I’ll add installation and usage instructions here.

**Important:** verifying group axioms (associativity, identity, inverses) and subgroup properties (e.g. normality), etc., is left to the user. The focus of this project is on performing group-theoretic computations that would be tedious by hand, not on theorem or property verification (which is better suited to systems like Lean).

## Dependencies

The core libraries do not use any external dependencies.

Once the CLI incorporates SQL (Postgres + migration / query tooling), those dependencies will be listed here.

## Roadmap

Due to the large potential scope of this project, this roadmap is not fixed and not necessarily in implementation order. Items may be split further as the project grows. In particular, there is certainly more planned for the future than what is displayed below

### Core Structures

- [X] Group Inteface
- [X] Group Presentations
  - [X] Word Sets Implementation
  - [X] Free groups
  - [X] Abelian Group Presentations
  - [ ] Abelian Group Detection
  - [ ] Dehn Presentation Detection
  - [ ] Resdiual Finiteness
  - [X] Classification of Presentations
    - [ ] Presentation Metadata
  - [ ] Tietze Transformations
- [ ] Finite Groups
  - [ ] Cayley Tables
  - [ ] Permutation Groups
- [ ] Automatic groups
- [ ] Varieties of groups
- [ ] Subgroups and Quotients

### Algorithms

- [X] Word Free and Cyclic Reductions
- [X] Subword Detection
- [ ] Multiplier Automata
- [ ] Word Problem Solvers:
  - [X] Free Groups
  - [X] Cyclic Groups
  - [X] Free Abelian Groups
  - [ ] Abelian Groups
  - [ ] Dehn Presentations
  - [ ] One-Relator Groups
  - [ ] Partial solution for the general case
- [ ] Normal Form Computations
- [ ] Conjugacy Problem Solvers

### Operations on Groups

- [ ] Direct Product
  - [ ] Semidirect Product
- [ ] Free Product
- [ ] Amalgams
- [ ] HNN Extensions
- [ ] Homomorphisms
  - [ ] Kernel Computations
  - [ ] Image Computations
  - [ ] Automorphisms

### CLI tool

- [ ] Pretty-printing
- [ ] SQL‑backed storage (Postgres)

## Contributions

This project is currently a personal exploration, and I’m **not accepting external pull requests** for now. Once the API stabilizes, this may change.

However, if you spot a bug, please open an issue on GitHub.

Suggestions that stay strictly within group theory are most welcome. For example, a request to add matrix groups over a given ring may fit, but I don’t plan to implement substantial ring theory here (that would belong in a separate project!).

