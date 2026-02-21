# 🚀 DevTrack

**DevTrack** is a simple, lightweight command-line tool built in Go to track how much time you spend on your daily development tasks. 

---

## ✨ Features

- Track tasks with simple start and stop commands.
- Organize task by projects.
- View detailed session history with list.
- Check active tasks details.
- Get daily and weekly productivity stats.
- Data stays on your machine (SQLite).
---

## 🛠 Installation

### Build from source

```bash
git clone https://github.com/rahul-nakum14/devtrack.git
cd devtrack
go build -o devtrack main.go
```

To run `devtrack` from any directory, move the binary to your local bin:
```bash
sudo mv devtrack /usr/local/bin/
```

---

##  Usage

###  Start tracking a task
Start a new session by providing a task name. Use the `-p` flag for a project name (optional).

```bash
devtrack start "feature-ui" -p "myproject"

# Output:

Started tracking: feature-ui
```

###  Stop the active session
Stops the currently running task and records the end time.

```bash
devtrack stop

# Output:

topped tracking: feature-ui
```

### List Active Sessions
Quickly see which task is currently being tracked.

```bash
devtrack active

# Output:
Seesions Are feature-ui
```

###  View  stats
Get a summary of time spent on each task today.

```bash
devtrack stats today

# Output:
Stats for today:
feature-ui 45m 20s
bug-fix    12m 10s
----------------
Total      57m 30s
```

```bash
devtrack stats week

# Output:
Stats for last week day:
feature-ui 45m 20s
bug-fix    1m 10s
----------------
Total      5m 30s
```

---

##  Available Commands

| Command | Description |
| :--- | :--- |
| `start <task>` | Start tracking a new task |
| `stop` | Stop the current active session |
| `active` | Show the currently running task |
| `list` | Show a history of all recorded sessions |
| `stats today` | View total time spent on tasks today |
| `stats week` | View time spent over the last 7 days |

---

##  Data Storage

DevTrack stores all your data locally in a SQLite database

```bash
~/.devtrack/devtrack.db
```

---

## Project Structure

- `cmd/devtrack/` - CLI command logic (Cobra)
- `internal/db/` - SQLite database initialization
- `internal/repository/` - Database queries and data handling
- `internal/service/` - Core business logic
- `internal/model/` - Data structure definitions
