## Functional Requirements

1. Create task
1. View task list
1. Update task
1. Complete task
1. Delete task

## Core Entities

- Task
  - Id: integer
  - Name: string
  - CreatedDate: datetime
  - DueDate: datetime
  - Priority: number
  - State: [ NotStarted, InProgress, Completed, Abandoned ]

## Commands

- `add` command
  - `todocli add <name>`
- `list` command
  - `todocli list`
- `edit` command
  - `todocli edit <id> <name>`
- `remove` command
  - `todocli remove <id>`