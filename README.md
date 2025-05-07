<p align="center">
  <img src="/docs/logo.png" alt="ansiwalker logo" width="525"/>
</p>

`ansiwalker` is an extremely lightweight Go package for iterating through strings or byte slices while skipping all ANSI/VT100 escape sequences, ensuring that only printable runes are processed.

## ✨ **Features**

* ANSI‑aware iteration: skips CSI, OSC, DCS, SOS, PM, APC, and C1 control codes.
* UTF‑8 safe: decodes and returns full runes for proper Unicode support.
* Simple API: single function to advance to the next visible character.
* Lightweight and dependency‑free.

## 🚀 **Getting Started**

```bash
go get github.com/galactixx/ansiwalker@latest
```

## 📚 **Usage**

```go
import (
    "fmt"
    "github.com/galactixx/ansiwalker"
)

func main() {
    input := "\x1b[31mError:\x1b[0m Something went wrong"
    buf := strings.Builder{}
    i := 0
    for i < len(input) {
        r, size, next, ok := ansiwalker.ANSIWalk(input, i)
        if ok {
          buf.WriteRune(r)
        }
        i = next
    }
    fmt.Println(buf.String())
}
```

### Output:
```text
Error: Something went wrong
```

## 🔍 **API**

`ANSIWalk(s string, i int) (r rune, size int, next int, ok bool)`
  Skips any ANSI escape sequence starting at index `i` and returns:

  * `r`: the next visible rune (or `0` at EOF).
  * `size`: the byte length of `r` in UTF‑8.
  * `next`: the index to resume parsing from.
  * `ok`: `false` if end of string is reached.

## 🤝 **License**

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## 📞 **Contact**

For questions or support, please open an issue on the [GitHub repository](https://github.com/galactixx/ansiwalker/issues).
