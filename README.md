# Anagram Finder (Go)

This project provides a simple tool to **find all sets of anagrams** in a given dictionary of words.  
It supports Unicode (e.g., Russian words) and ignores duplicates.  

---

## Project Structure

```

.
├── cmd/anagram/          # main.go — entry point
├── internal/anagram/     # package with logic and tests
│   ├── finder.go
│   └── finder_test.go
├── go.mod
├── Makefile
└── README.md

````

---

## How it Works

The `FindAnagrams` function takes a slice of strings (words) and returns a map where:

* **Key:** the first encountered word in each anagram set  
* **Value:** a slice of all words that are anagrams of each other, sorted in ascending order  

Words with no anagrams are **excluded**.  
All words are normalized to lowercase, and duplicate words are ignored.

**Example:**

Input:

```go
[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
````

Output (keys order may vary):

```go
"пятак": ["пятак", "пятка", "тяпка"]
"листок": ["листок", "слиток", "столик"]
```

---

## Installation

Clone the repository:

```bash
git clone https://github.com/aliskhannn/anagram.git
cd anagram
```

---

## Build

```bash
make build
```

Binary will appear in `bin/anagram`.

---

## Usage

```bash
./bin/anagram
```

---

## Testing

Run unit tests:

```bash
make test
```

---

## Linting

```bash
make lint
```

Uses:

* `go vet`
* `golangci-lint`

---

## Clean

```bash
make clean
```