# Building manually

## Pre-requisites
You need golang installed to build this. I tend to use very new versions since vulnabilities move fast in the underlying libraries, so I'd recommend visiting the [golang download page](https://golang.org/dl/) and getting the latest version. At the time of writing this, I'm far ahead of the golang version included in the debian/ubuntu repos.

You'll also need `make` installed, and libc6 (which should be installed - just a note in case you're trying to run on alpine)
Debian/Ubuntu:
```bash
sudo apt install make libc6
```
`libc6` is essential so that the sqlite stuff works. I haven't done much testing with other libraries, ie, musl, so I can't guarantee it'll work with compat packages.

## Building

Now you can just run `make` in the root of the repo. This will build a binary called `wmb` in the same folder.

# Building with docker
Standard docker build command:
```bash
docker build -t wmb:latest .
```

# Running with debugger in VSCode
There's a sample `launch.json` in the `.vscode` directory. Rename it, fill in some enviornment variables, and you should be good to go.

# Notes
I publish docker images of the `main` branch on github. You can pull them from `ghcr.io/cfindlayisme/wmb:latest`

Builds are done in `amd64` and `arm64` by the pipeline. If you need a different architecture, you'll have to build it yourself.