# Golang Webassembly Degreed to Radians converter

You can see the result [here](https://amirmahdyjebreily.github.io/WasmDegreeToRadians/) !

<p>This project is just an experiment to use golang in the web environment and use MathJax to display
    mathematical and mathematical expressions. You can see all the technologies that were combined in this
    project:</p>
<center>
<a href="https://www.mathjax.org">
    <img title="Powered by MathJax" src="https://www.mathjax.org/badge/mj_logo.png" border="0"
        alt="Powered by MathJax" />
</a>
<a href="https://webassembly.org/">
    <img title="Powered by Wasm" src="https://upload.wikimedia.org/wikipedia/commons/thumb/1/1f/WebAssembly_Logo.svg/240px-WebAssembly_Logo.svg.png" width="72" height="72" border="0"
        alt="Powered by Wasm" />
</a>
<a href="https://webassembly.org/">
    <img title="Powered by Golang" src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" width="96" height="96" border="0"
        alt="Powered by Golang" />
</a>
</center>

## Build and Usage

1. ### Clone the Repo :
   ```shell
   git clone https://github.com/AmirMahdyJebreily/WasmDegreeToRadians.git
   ```
   go to the `WasmDegreeToRadians` dirctory with the fllowing command:
   ```shell
   cd WasmDegreeToRadians
   ```
2. ### Build Webassembly

   Unix terminals:

   ```shell
   GOOS=js GOARCH=wasm go build -o ./out/radian.wasm
   ```

   Windows command propmt:

   ```shell
   $Env:GOOS = "js"; $Env:GOARCH = "wasm"; go build -o ./out/radian.wasm
   ```

3. ### Use a file server

   you can create a http file server using golang like this :

   ```go
    package main
    import (
        "fmt"
        "net/http"
    )
    func main() {

        err := http.ListenAndServe(":9090", http.FileServer(http.Dir("../../assets")))

        if err != nil {
        	fmt.Println("Failed to start server", err)
    	    return
        }

    }
   ```  
   or you can use [this practical vscode extention](https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer) to host yout html files in local network.

___
Star this repo ⭐ and register an issue if you see a problem. Thank you for reading  
Code Agha ❤️
