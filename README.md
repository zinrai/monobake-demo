# monobake-demo

A demo of building Docker images with GitHub Actions and [monobake](https://github.com/yourorg/monobake).

## Applications

| App      | Description      | Port |
|----------|------------------|------|
| frontend | Returns HTML     | 8080 |
| backend  | Returns JSON     | 8000 |
| worker   | Prints to stdout | -    |

## Local Build

See [build-demo.sh](build-demo.sh).

## Release with monobake

Tag a release

```bash
$ git tag backend/v0.1.0
$ git push origin backend/v0.1.0
```
