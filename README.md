asana-go
========
[![wercker status](https://app.wercker.com/status/e7088601a3a1a737eeaa9f47a960648f/m "wercker status")](https://app.wercker.com/project/bykey/e7088601a3a1a737eeaa9f47a960648f)

Asana Client for Go

## SYNOPSYS
```go
import (
  "github.com/HirokazuMiyaji/asana-go"
)

a := asana.New()

# User
a.Users().Me()
a.Users().All()

userId := 1
a.Users().Where(userId)
```
