# sdkgen

[![Coverage Status](https://coveralls.io/repos/github/alokic/sdkgen/badge.svg?t=4uMFW5)](https://coveralls.io/github/alokic/sdkgen)
[![Build Status](https://drone.alokic.com/api/badges/alokic/sdkgen/status.svg)](https://drone.alokic.com/alokic/sdkgen)

Generate `openapi 3 API spec` from `json` request and responses of `REST APIs`.

## Technical Specs

  https://alokic.atlassian.net/wiki/spaces/EN/pages/556957804/Roles+and+Permissions

## Health Check

Standard health check end point: `/report/health`

## Local Setup

### Pre-requisites

```md
# install go
$ brew install go@1.11

# install dep
$ brew install dep

```

### Build

1. Make sure GOROOT and GOPATH are properly set.
2. Set PATH

    ```export PATH=$PATH:$GOPATH/bin```

3. Git clone the repo in $GOPATH/src/github.com/alokic.

    ```cd $GOPATH/src/github.com/alokic```

    ```git clone git@github.com:alokic/sdkgen.git```

    ```cd sdkgen```

4. Check all deps are there by running below

   ```$dep ensure```

5. Export required environment variables


6. Build the app

   ```make build```

   ```'sdkgen' gets installed in $GOPATH/bin```

7. Help

   ```sdkgen --help```

8. Run the app

   ```sdkgen openapi --in <input_path> --out <output_path>``` # generates only `openapi` specs

9. Output

Following is the structure of output folder.
`SDKGEN_OUTPUT_PATH`

### tests

1. Run all tests now

  ```make test```

## Openapi generator commands

### Config options for a language

```md
openapi-generator  config-help -g python
```

### Generate code

```md
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
