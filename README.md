# ⚠️ Disclaimer
This tool is for educational and research purposes only. The author is not responsible for any misuse or illegal activity.

# godstoken

A powerful tool for checking the validity of Discord tokens.  
Supports single and batch checks, optional user info scraping, and debug logging.

## ✨ Features

- ✅ Check a single token with `-t`
- 📁 Check a list of tokens from a file with `-f`
- 🔍 Collect basic user info (`-d`) like username
- 🐛 Enable debug logging to see each action step-by-step (`-debug`)
- ⚠️ Automatic argument validation (prevents `-t` and `-f` being used together)

## 📦 Usage

```bash
godstoken [arg]
```

### Options

| Flag         | Description                                |
|--------------|--------------------------------------------|
| `-t=""`       | Check a single token                      |
| `-f=""`       | Check tokens from a file (one per line)   |
| `-d`          | Collect basic user information            |
| `-debug`      | Enable debug logging                      |

### ❗ Conflict

You **cannot** use `-t` and `-f` at the same time.  
If both are passed, the program will throw an error.

## 🔧 Example

```bash
godstoken -t="your_token_here"
godstoken -f="tokens.txt"
godstoken -t="your_token_here" -d -debug
```

## 📁 Project Structure

```
godstoken/
│
├── main.go                 # entry point
├── go.mod                  # go module
└── pkg/
    ├── file/               # file reader
    │   └── read.go
    ├── token/              # token checker
    │   └── check.go
    ├── user/               # user info collector
    │   └── fetch.go
    └── util/               # debug logger
        └── debug.go
```

---
🔥 Use responsibly — no monkey business.

