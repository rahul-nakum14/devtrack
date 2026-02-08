# DevTrack

DevTrack is a simple command‑line tool to track how much time you spend on development tasks.
It works completely offline and stores data locally using SQLite.

---

## Features

- Start and stop work sessions
- Prevent overlapping sessions
- Track tasks with optional project names
- View today’s work summary
- Local SQLite database (no server required)

---

## Installation

### Build from source

```bash
git clone https://github.com/rahul-nakum14/devtrack.git
cd devtrack
go build -o devtrack
```

## Usage

### Start tracking a task
```bash
./devtrack start <task>

Example:
./devtrack start coding

Also works with project name (optional):
./devtrack start bugs-fixes --project demoproject
```

### Stop the active session
```bash
./devtrack stop

This stops the currently running session and records the end time.
```
### View today’s stats

```bash
./devtrack stats today

Stats for today:
coding     0m 31s
testing    11m 58s
----------------
Total      12m 29s
```
---

### Data Storage

DevTrack stores all data locally using SQLite.

Database file location:
```bash
~/.devtrack/devtrack.db
```

## Available Commands

| Command | Description |
|--------|------------|
| `start <task>` | Start tracking a task |
| `stop` | Stop the active session |
| `stats today` | Show today’s work summary |


