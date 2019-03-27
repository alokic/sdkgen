# sdkgen

[![Coverage Status](https://coveralls.io/repos/github/alokic/sdkgen/badge.svg?t=4uMFW5)](https://coveralls.io/github/alokic/sdkgen)
[![Build Status](https://drone.alokic.com/api/badges/alokic/sdkgen/status.svg)](https://drone.alokic.com/alokic/sdkgen)

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

   In following example, `openapi` specs for `jsons` in `./examples/input` are generated in `./examples/out`.

      ```bash
      docker run -v ${PWD}/examples:/examples sdkgen  openapi -i /examples/input -o /examples/output/openapi
      ```

4. Generates SDK

   In following example `sdk` in `python` is generated in `./examples/output/sdk/python` directory from `./examples/output/openapi/main.yaml` generated in step 7.

      ```bash
       docker run -v ${PWD}/examples:/examples sdkgen  sdk -i /examples/output/openapi/main.yaml -o /examples/output/sdk -g python
      ```

   In following example, `sdk` is generated with `configs` defined in `./examples/config/python.json`.

      ```bash
       docker run -v ${PWD}/examples:/examples sdkgen  sdk -i /examples/output/openapi/main.yaml -o /examples/output/sdk -g python -c /examples/config/python.json
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

   In following example, `openapi` specs for `jsons` in `./examples/input` are generated in `./examples/output/openapi`.

      ```bash
      ctl openapi -i ./examples/input -o ./examples/output/openapi
      ```

8. Generates SDK

   In following example `sdk` in `python` is generated in `./examples/sdk/python` directory from `./examples/output/main.yaml` generated in step 7.

      ```bash
      ctl sdk -i ./examples/output/openapi/main.yaml -o ./examples/output/sdk -g python
      ```

   In following example, `sdk` is generated with `configs` defined in `.examples/config/python`.

      ```bash
      ctl sdk -i ./examples/output/openapi/main.yaml -o ./examples/output/sdk -g python -c ./examples/config/python.json
      ```

### tests

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

## Release Process

We follow [Feature branch workflow](https://www.atlassian.com/git/tutorials/comparing-workflows/feature-branch-workflow)

Feature branches are merged in `master` branch after `PR` approval.

Deployment is triggered when we tag `master` branch.

Tagging convention for `production` environment is `v.MAJOR_NUMBER.MINOR_NUMBER.PATCH_NUMBER`.

We increment `NUMBERS` when:

* `MAJOR_NUMBER` when breaking backward compatibility.
* `MAJOR_NUMBER` when adding a new feature which doesn’t break compatibility.
* `PATCH_NUMBER` when fixing a bug without breaking compatibility.

Tagging convention for `staging` is `v.MAJOR_NUMBER.MINOR_NUMBER.PATCH_NUMBER-<SUFFIX>.<CHILD_SUFFIX>`.

* `SUFFIX` is mandatory for `staging` deployments. Valid `SUFFIX` are `alpha`, `beta`, `rc`.
* `CHILD_SUFFIX` is optional and should be used for incremental updates in a tag like `v.1.0.0-rc.1`
