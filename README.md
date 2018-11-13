# Noisefloor software synth

Module software Synthesizer for Javascript and native code targets

## Requirements

1. Linux Build environemnt

   This is requred for the gopherjs transpiler - building under Windows is posible with `Bash on Ubuntu on Windows`

2. `go`

   Only Go 1.11 has been tested, other versions may work    

3. Ensure go bin directory in path, for example
   `export PATH=$GOPATH/bin:$GOROOT/bin:$PATH`

   this is needed for `dep` and `gopherjs`

4. `git` which `go get` command will use to install go packages

5. Go `dep` for package management/version control   
   https://github.com/golang/dep

6. A web browser to run the transpiled javascript in

## Build instructions

Only Javascript build target is currently available.

To Build

2. Get package with go Get

   `go get github.com/jacoblister/noisefloor`

3. change to directory

   `cd ~/go/src/github.com/jacoblister/noisefloor`

4. Install dependencies with `dep`

   `dep ensure`

5. build and install `gopherjs` (javascript transpiler modified with optimisations)

   `(cd vendor/github.com/gopherjs/gopherjs && go install)`

6. build the project

   `cd ~/go/src/github.com/jacoblister/noisefloor/js/noisefloor`
   `gopherjs build`

7. Open in a web browser

   `/[yourgopath]/src/github.com/jacoblister/noisefloor/js/noisefloor/index.html`
