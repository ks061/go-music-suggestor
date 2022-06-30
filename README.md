1. **Report uploaded on Moodle and can be accessed in the `main.go` file.**

<hr/>

2. The following are answers to the **software questions**:

**What your software does:** The program prompts the user for how they are feeling (from a selection of the following emotions: happy,
despair, peaceful, and angry). It then selects music that emulates how the user is feeling. _Note: The detailed software description is at the top of the `songs.go` file with each struct and function thoroughly documented._

**Why it is appropriate to do this in your language (1 paragraph):** Go has great concurrency for preprocessing data in data mining and machine learning applications through its built-in `sync` library, including the `WaitGroup` construct. It also offers the `go` and `defer` commands for easy coordination of parallel processing sandwiched between consecutive processing. Go is also strongly typed, allowing the programmer to catch more errors at compile time. At the same time, it is also implicitly typed, a convenience that enables the programmer to quickly set and use short-lived variables without explicitly writing out its type each time. Go effectively decouples the notion of being strongly typed, which allows for robust error-checking and needing to explicitly type variables, which can sometimes hinder agile rapid code development. Another interesting convenience is that while loops are written as pared-down for loops, so to speak, in Go, which allows the user to consolidate their programming loops into for loops.

**A concise (short) user manual to let us know how to compile, run, and use your software (up to a page is fine):**
1. The software is written in `songs.go`.
2. Assuming go is installed on the command line; if not, follow these instructions to install go: https://go.dev/doc/install.
3. `cd` into the directory with the `songs.go` file.
4. Enter `go run songs.go` into the command prompt

<hr/>

3. **Where is the video? See `demo.mov` in this repository.**


