# Stack

Push and pop tasks in this TUI Kanban app.

## Database

Data is stored locally.  In a idempotent JSONL WAL.

```jsonc
// schema
// Id and timestamp are required.  Other fields provided as/when they change.
{
    "id": "guid",
    "timestamp": "unix-time",
    "title": "title",
    "body": "markdown-body",
    "addedTags": [
        // 0, 1, or n.
        "a-z0-9-"
    ],
    "removedTags": [
        // 0, 1, or n.
        "a-z0-9-"
    ]
}
```
