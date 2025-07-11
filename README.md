# GoLang Todo CLI

A simple command-line task management system written in Go.

## Features

- View all current tasks
- Add new tasks
- Remove completed tasks
- Simple command-based interface

## Usage

1. **Run the application:**

```sh
   go run main.go
```

## Available Commands:

- Tasks — List all current tasks
- Add — Add a new task
- Remove — Remove a task by its number
- Exit() — Exit the program

## Example Session:

```md
Welcome To Your Task Management System!
Please input 'Tasks' to see any current Tasks
Please input 'Add' to add a new Task
Please input 'Remove' to remove a completed Task
Please input 'Exit()' to exit the program

Input New Command:
add
What is your new task?
Buy groceries
Task Added!

Input New Command:
tasks
Tasks To do:
  1. Buy 
```
## Roadmap
Upcoming Feature: Persistent Local Storage
Currently, tasks exist only in memory and are lost when the program exits.
Upcoming Feature: Tasks in JSON-Struct format and with timestamps,priority levels, and groupings
Planned enhancement:

Store tasks in a local file (e.g., tasks.txt or JSON) so that your to-do list is saved between sessions.
On startup, the app will load existing tasks from the file.
When adding or removing tasks, the file will be updated automatically.
Stay tuned for updates!