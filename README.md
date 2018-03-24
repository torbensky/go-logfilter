# go-logfilter
A Go (Golang) library for configuring log output levels at a package or file level.

# Usage

[Working example](./example/myapp.go)

```
func main(){
    gfl := logfilter.NewLogFilter()

    // Only log error or higher for a package
    glf.SetLevel(log.ErrorLevel, "github.com/my/package") 

    // Output debug or higher for a specific file. File matches are given priority over package matches.
    glf.SetLevel(log.DebugLevel, "github.com/my/package/file.go")

    // Load a bunch of filters from a config
    config := `
		file1.go:debug,
		file2.go:warn,
		github.com/torbensky/logfilter:panic
	`
    gfl = LoadConfig(config)

    // Wrap a logrus Hook with the filter
    hook := airbrake.NewHook(123, "xyz", "development")
    hook = logfilter.NewHookFilter(hook, gfl) // Now only the unfiltered entries will go to airbrake
}

```
