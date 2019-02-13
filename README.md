# fipsvc

[fipsvc](https://fipress.org/project/fipsvc) is to setup a service in Go, and help you with reload and clean up. 

**Usage**

1. Add cleanup and reload hooks
```
fipsvc.AddCleanupHooks(cleanup)
fipsvc.AddReloadHooks(load)
```


2. Start
```
fipsvc.Start()
```