# Noisefloor software synth

Module software Synthesizer for Javascript and native code targets

## Requirements

1. Linux Build environment

   This is required for the gopherjs transpiler - building under Windows is possible with `Bash on Ubuntu on Windows`

2. [`go`](https://golang.org)

3. Ensure go bin directory in path, for example

   ```
   export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
   ```

   this is needed for `gopherjs`

4. [`git`](https://git-scm.com/) which `go get` command will use to install go packages

5. [`gopherjs`](https://github.com/gopherjs/gopherjs) which is used for transpiling to js
   ```
   go get -u github.com/gopherjs/gopherjs
   ```

7. A web browser to run the transpiled javascript in

## Build instructions

Only Javascript build target is currently available.

To Build

1. Clone repository

   ```
   git clone https://github.com/jacoblister/noisefloor.git
   ```
   or 
   ```
   git clone git@github.com:jacoblister/noisefloor.git
   ```

2. Build the project

   ```
   cd noisefloor/build/js/noisefloor
   gopherjs build
   ```

3. Run the build
   ```
   gopherjs server
   ```
   Then open http://localhost:8080 in browser
