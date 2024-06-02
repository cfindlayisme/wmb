# Building manually
To build the image, just run `make` in the root of the repo. This will build a binary called `wmb` at the root of repository.

# Building with docker
Standard docker build command:
```bash
docker build -t wmb:latest .
```

# Running with debugger in VSCode
There's a sample `launch.json` in the `.vscode` directory. Rename it, fill in some enviornment variables, and you should be good to go.

# Notes
I publish images of the `main` branch on github. You can pull them from `ghcr.io/cfindlayisme/wmb:latest`