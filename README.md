# LangRoutine
> Asynchronous analysis of text in files to determine language. Based on lingua-go library

The project is designed to demonstrate working with goroutines, canals. Also such synchronization primitive as WaitGroup was used.

### Requirements:
![golang](https://badgen.net/static/go/1.22.0/green?icon=github)<br/>
You can install Golang <a href="https://go.dev/doc/install">there</a>

### Installing:
1. Clone repository 
2. In main directory:<br/>
  ```bash
  go build\

  //for windows
  backendproj.exe

  //for linux
  ./backendproj
  ```

### Usage:
  The program can work in synchronous and asynchronous mode (depending on the flag). The program can also read not files in the directory, but a specific file.
  Flags:
  ```
  -fpath string
        file path
  -path string
        folder path (default "assets")
  -sync
        choose sync mode
  ```
