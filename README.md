# gofilelogger
A Go (Golang) library for configuring log output levels at a package or file level.

# Usage

```
func main(){
    gfl := gofilelogger.New()

    // Only log error or higher for a package
    glf.SetLevel(log.ErrorLevel, "github.com/my/package") 

    // Output debug or higher for a specific file. File matches are given priority over package matches.
    glf.SetLevel(log.DebugLevel, "github.com/my/package/file.go") 
}

```