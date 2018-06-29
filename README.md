# kill-port

[![MIT License][license-badge]][LICENSE]
[![PRs Welcome][prs-badge]][prs]

This package provides a powerful cross function to kill (Windows / MacOS / Linux) process running on given port.

## Installation

```
go get github.com/rimiti/kill-port
```

## Examples

```
# Kill process running on port 80 (HTTP)
$ kill-port 80

# Kill process running on port 443 (HTTPS)
$ kill-port 443

# Kill process running on port 3306 (MySQL)
$ kill-port 3306
```

## License

MIT © [Dimitri DO BAIRRO](https://github.com/rimiti/kill-port/blob/master/LICENSE)

[license-badge]: https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square
[license]: https://github.com/rimiti/kill-port/blob/master/LICENSE
[prs-badge]: https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square
[prs]: http://makeapullrequest.com
