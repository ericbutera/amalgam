# Code Quality

Be sure to install the pre-commit hooks which run various linters, formatters, and tests.

```sh
just setup
```

## Linters

- [golangci-lint](https://golangci-lint.run/) - golang linters (config in [.golangci.yml](../.golangci.yaml))
- [eslint](https://eslint.org/) - typescript linters
- [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) - enforced by pre-commit hooks
