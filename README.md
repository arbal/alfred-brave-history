# alfred-brave-history

Search Brave history from Alfred and open in browser.

## Features

- Search Brave history (title and URL)
- Support another Brave profile

## Installation

Clone and `make dist` or just download [binary releases](https://github.com/pasela/alfred-brave-history/releases).

```sh
git clone https://github.com/pasela/alfred-brave-history.git
cd alfred-brave-history
make dist
open alfred-brave-history.alfredworkflow
```

## Usage

in Alfred:

```
ch {query}
```

## Use another Brave profile

1. Open workflow `Brave History` in Alfred Workflows tab.
2. Open Workflow Configuration dialog by upper right side button.
3. Set `BRAVE_PROFILE` variable with your Brave profile directory name or path such as `Profile 1`.

## License

MIT

## Author

Yuki (a.k.a. pasela)
