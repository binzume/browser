# Minimum web browser implementation

Just a plain web browser implemented in Golang.

## Usage

### Install

```bash
go install github.com/binzume/browser/cmd/browser
```

### Accees Google!

```bash
browser get https://www.google.com/
```

Nothing is displayed?, but it is okey. This browser does not have the ability to render HTML.

### Golang interface

```go
import 	"github.com/binzume/browser"

client, _ := browser.NewWebBrowser()
req, _ := http.NewRequest("GET", os.Args[2], nil)
response, _ := client.Do(req)
```

`WebBrowser` is `http.Client`.

## License

MIT License
