# Development

## Tooling

This project uses [Magefile](https://magefile.org/) to simplify some tasks:

- use `mage docs` to spin up a mkdocs server at `http://localhost:8000`
- use `mage proto` to compile all proto files in `internal/proto`
- use `mage clean` to remove all compiled proto files from `pkg/pb`