---
title: Graceful Shutdown
menu:
  side:
    parent: recipes
    weight: 13
---

### With [grace](https://github.com/facebookgo/grace)

`server.go`

{{< embed "graceful-shutdown/grace/server.go" >}}

### With [graceful](https://github.com/tylerb/graceful)

`server.go`

{{< embed "graceful-shutdown/graceful/server.go" >}}

### Maintainers

- [mertenvg](https://github.com/mertenvg)

### Source Code

[graceful](https://github.com/labstack/echo/blob/master/recipes/graceful-shutdown/graceful)

[grace](https://github.com/labstack/echo/blob/master/recipes/graceful-shutdown/grace)
