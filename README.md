# Mathematical Groups in Go

An implementation of mathematical groups in Go, designed by a geometric group theorist. This project is a personal exploration and not connected to any academic program.

## Motivation

During my master's degree in mathematics, I specialized in geometric group theory. Amazed by the deep links between theoretical computer science and geometry, I decided to build this bridge between the beautiful theory of groups and programming. With the arrival of generics in Go 1.18 (2022), it finally became realistic to model mathematical structures in a way that feels natural. This project is my attempt to bring the world of group theory into a modern, strongly typed language, in a style that could plausibly fit a "math/group" package in the Go standard library. I hope this project will encourage more practical applications of group theory; as groups arise naturally in many contexts, there certainly is plenty of untapped potential!

Bridging programming and group theory is highly non-trivial, and as new research discoveries are found, this project will reflect those and thus will never truly be "finished". Mathematical proofs describing an algorithm (or proving the mere existence of such an algorithm) do not establish the careful algorithmic and architectural design required to implement in software. Even a single step in a mathematical proof, like word reduction or finding subwords, requires careful considerations, such as in time complexity. Other algorithms require initiating multiple processes concurrently, another reason I chose Go for this project. Therefore, one aim of this project is to translate those theoretical procedures into efficient implementations suitable for real computation.

## Tutorial

This project is a Go library, hence the first step is to import this repository's `presentation` package. We will take the group $G = \mathbb{Z} \times \mathbb{Z} / 2 \mathbb{Z}$ as an example, which has a presentation $\langle a , b | aba^{-1}b^{-1}, b^2 \rangle$.

We start by constructing the relations. In this library, words are defined using generator exponent pairs; the first generator is 0, the second generator is 1, and so on.

```go
abComm := NewWord({{0,1},{1,1},{0,-1},{1,-1}})
bSquare := NewWord({{1,2}})
```
Note: A RawWord type also exists, which carries less encoding data. However, it is not recommended to use them because they cannot be elements of the `WordSet` type.

Then we define the set of relations using a `WordSet`:

```go
rels := make(WordSet)
rels.add(abComm)
rels.add(bSquare)
```
Alternatively, we could have also done `rels := NewWordSet([]Word{abComm, bSquare})` instead.

Now, we have everything necessary to define the presentation above:

```go
G, err := NewGroupPresentation(2, rels)
```
Here, the 2 corresponds to the number of generators and `rels` to the WordSet of relations. It is important that $G$ here is a pointer to the `GroupPresentation` struct, so that we may mutate it with methods, up to isomorphism. In particular, Tietze transformations are planned in the future. Furthermore, the error is not `nil` if the relation WordSet involves more generators than stated.

Every presentation has a `classes` field (e.g. `Trivial`, `Free`), which are listed in the presentation/classes.go file. The function `NewGroupPresentation` automatically adds the relevant classes (and their negation) if the presentation is cyclic or has only one relator. Otherwise, we add the classes manually. In this case:

```go
G.AddClass(Abelian, true)
G.AddClass(Trivial, false)
```
The boolean in the `AddClass` determines if the group presentation is part of the class or not. In this instance, we use `true` for `Abelian` but `false` for `Trivial` because $G$ is indeed abelian but not a presentation of the trivial group.

## Current Scope

This project is currently a **library**, not a standalone tool. A CLI tool is planned for the future; once it exists, I’ll add installation and usage instructions in the tutorial section.

**Important:** verifying group axioms (associativity, identity, inverses) and subgroup properties (e.g. normality), etc., is left to the user. The focus of this project is on performing group-theoretic computations that would be tedious by hand, not on theorem or property verification (which is better suited to systems like Lean).

So far, the project has focused on group presentations and the presentations package is the one most ready to be used. All functionality in documented in the files themselves. The word.go file describes how words are represented in the library and comes with many useful operations like word reduction and subword-finding.   word_encode.go describes the canonical string representation of words that the library uses. Right now, this canonical string representation is mainly used as keys for WordSet maps, which are defined in wordset.go. As a result, the library's Word struct has the string representation as one of its fields, and it is recommended to use Word instead of RawWord at all times.

Group presentations are defined in presentation.go and utilize WordSets for the set of relations to emulate set-behavior instead of slice behavior. Each presentation has a classes field that classifies the properties the group has (for example cyclic or abelian). The full list of classes are listed in classes.go, along with functions to manually add and remove classes from a presentation. Specialized presentation constructors like those found in free.go and abelian.go are available for certain of these classes. Finally, reduce.go features word problem solutions for certain classes of groups via reduction of words to a normal form.

The remaining files are either test files, unfinished, or specific utilities for other functions in the library, and thus of not much interest to users yet. For instance, utils.go and kmp.go fall in the latter category.

## Dependencies

The core libraries do not use any external dependencies.

Once the CLI incorporates SQL (Postgres + migration / query tooling), those dependencies will be listed here.

## Roadmap

Due to the large potential scope of this project, this roadmap is not fixed and not necessarily in implementation order. Items may be split further as the project grows. Moreover, new features will be added to the roadmap as the project grows.

NB: I am currently improving the implementations of words using a tree structure instead of a slice of arrays. 

### Core Structures

- [X] Group Interface
- [X] Group Presentations
  - [X] Word Sets Implementation
  - [X] Free groups
  - [X] Abelian Group Presentations
  - [ ] Abelian Group Detection
  - [ ] Dehn Presentation Detection
  - [ ] Residual Finiteness
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

- [X] Free and Cyclic Reductions
- [X] Subword Detection
 - [X] Using KMP
 - [X] Primitive Word Roots
- [ ] Rewriting Systems
- [ ] Multiplier Automata
- [ ] Word Problem Solvers:
  - [X] Free Groups
  - [X] Cyclic Groups
  - [X] Free Abelian Groups
  - [ ] Abelian Groups
  - [ ] Dehn Presentations
  - [ ] One-Relator Groups
  - [ ] Residually Finite Groups
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

