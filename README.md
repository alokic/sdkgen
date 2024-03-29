# sdkgen

Generate `openapi 3 spec` from `json` request and responses of `REST APIs`.

## Docker Setup

### Pre-requisites

#### Install docker

For macOS -> https://docs.docker.com/docker-for-mac/install/

### Build

1. Git clone the repo.

```bash

    mkdir -p go

    cd go

    git clone git@github.com:alokic/sdkgen.git

    cd sdkgen
```

2. Build docker image

```bash
    docker build -t sdkgen .
```

3. Generates `openapi` specs

   In following example, `openapi` specs for `jsons` in `./examples/inputs/nested` are generated in `/tmp/openapi`.

```bash
    docker run -v ${PWD}/examples:/examples -v /tmp:/tmp sdkgen  openapi -i /examples/inputs/nested -o /tmp/openapi
```

4. Generates SDK

   In following example `sdk` in `python` is generated in `/tmp/sdk/python` directory from `/tmp/openapi/main.yaml` generated in step 7.

```bash
    docker run -v ${PWD}/examples:/examples -v /tmp:/tmp sdkgen  sdk -i /tmp/openapi/main.yaml -o /tmp/sdk -g python
```

   In following example, `sdk` is generated with `configs` defined in `./examples/config/python.json`.

```bash
    docker run -v ${PWD}/examples:/examples sdkgen  sdk -i /tmp/openapi/main.yaml -o /tmp/sdk -g python -c /examples/config/python.json
```

## Local Setup

### Pre-requisites

```bash
    # Install go
    $ brew install go@1.12

    # Install dep (golang package manager)
    $ brew install dep

    # install openapi generator. Refer
    $ brew install openapi-generator
```

### Build

1. Make sure GOPATH is properly set.
If you are new to `golang` then,

```bash
    makdir -p $HOME/go 
    echo 'export GOPATH=$HOME/go' >>~/.bash_profile
    mkdir -p $GOPATH/src $GOPATH/pkg $GOPATH/bin
```

2. Set PATH

```bash
    export PATH=$PATH:$GOPATH/bin
```

3. Git clone the repo.

```bash
    mkdir -p $GOPATH/src/github.com/alokic

    cd $GOPATH/src/github.com/alokic

    git clone git@github.com:alokic/sdkgen.git

    cd sdkgen
```

4. Check all deps are there by running below

```bash
   dep ensure
```

5. Build the app

```bash
   make build
```

   `ctl` gets installed in $GOPATH/bin

6. Help

```bash
    ctl --help
```

7. Generates `openapi` specs

   In following example, `openapi` specs for `jsons` in `./examples/inputs/nested` are generated in `/tmp/openapi`.

```bash
    ctl openapi -i ./examples/inputs/nested -o /tmp/openapi
```

8. Generates SDK

   In following example `sdk` in `python` is generated in `./examples/sdk/python` directory from `/tmp/output/main.yaml` generated in step 7.

```bash
    ctl sdk -i /tmp/openapi/main.yaml -o /tmp/sdk -g python
```

   In following example, `sdk` is generated with `configs` defined in `.examples/config/python`.

```bash
    ctl sdk -i /tmp/openapi/main.yaml -o /tmp/sdk -g python -c ./examples/config/python.json
```

## Tests

```bash
    make test
```

## Openapi generator commands

For details check: [Openapi Codegen](https://github.com/OpenAPITools/openapi-generator)

### Config options for a language

```bash
    openapi-generator  config-help -g python
```

### Generate code

```bash
    openapi-generator generate   -i <input_file_path> -g python -o <output_folder> --enable-post-process-file
```
