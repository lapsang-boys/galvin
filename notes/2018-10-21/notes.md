# Galvin

Galvin is, yet to be, an esoteric programming language, a passion project and of
course a time zinc.

The main goal is to write fun code fast, in new paradigms.

## Concepts

From our language specification written in EBNF format we can use TextMapper to
generate our language parsing code.

The compiler will produce LLVM IR which can then be made into platform specific
machine code.

## Language features

- Named function call parameters (under investigation)
- Compile time fomal constraints (nice!)
- Runtime constraints (nice!)
- Type classes
- Atoms
- Sane defaults (=, and, not, or)
- Small keyword space
- must : Pattern matching, lol yah
- Lazy standard? prehaps :+1:
- Some form of IO, `do`, monad? chezksjobb


```galvin
contains(collection:xs, element:x)
xs.contains(element: x)
xs `contains` x

 -- This will create an anonymous curried function where the
named parameter collections will be filled as a positional argument.

xs.filter(contains(collection:_, element:x))

filter(predicate:even, functor:[1..10])

```

## Toolchain

The names of galvin tools and packages derive from metal and oxidization terms.
This stems from the etymology of Galvin, where in a discussion the word galvin
came up and we instantly thought it would be a good name for a programming
language referencing Rust and its antagonizer: galvanized steel.

Common naming patterns: gal-prefix, win-suffix, zinc related, metalworking pun[1, 2], or
referencing Luigi Galivani[3].

We could either go the rust route with stand-alone commands (e.g. `rustc`, `cargo`)
or sub-commands (e.g. `go fmt`, `go doc`).

For the sake of decluttering the binary-namespace, sub-commands feel superior.

Suggested names for integral tools in the toolchain, or project names for main
parts of the langauge:

__Compiler__
- `gal`

__Testing__
- `gal test`
- `gal win?` (Note: legal name in bash)

__Package manager__
- `gal pack`
- `anvil`

__Documentation module__
- `gal doc`

__Binding generator__
- `gal gen`
- `oxygen`

__Formatter__
- `gal fmt`
- `spangle` [4]

__Linter__
- `deburr` [5]

__Constraint engine__
- ?

__Formal proof engine__
- ?

__Type system__
- ?

__Garbage collection__
- `scrap`

## Package manager

A trivial dependency manager with central vendoring.

---

\[1\]: https://en.wikipedia.org/wiki/Category:Metalworking_terminology
\[2\]: https://en.wikipedia.org/wiki/Outline_of_metalworking
\[3\]: https://en.wikipedia.org/wiki/Luigi_Galvani
\[4\]: https://en.wikipedia.org/wiki/Galvanization#/media/File:Feuerverzinktes_Gel%C3%A4nder.jpg
\[5\]: https://en.wikipedia.org/wiki/Burr_(edge)
