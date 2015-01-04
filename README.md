asana-go
========
[![wercker status](https://app.wercker.com/status/e0a12d920eee78cafbe418dcf74abc21/m "wercker status")](https://app.wercker.com/project/bykey/e0a12d920eee78cafbe418dcf74abc21)

Asana Client for Go

## SYNOPSYS
```go
import (
  "github.com/hirokazumiyaji/asana-go"
)

a := asana.New()

# User
a.Users().Me()
a.Users().All()

userId := 1
a.Users().Where(userId)

// Task
a.Tasks().All()

taskId :=1
a.Tasks().Where(taskId)
```
