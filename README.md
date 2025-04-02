# php-cs-fixer-stdin

PHP CS Fixer wrapper for stdin text editor integration

## Context

Some code editors support stdin integration, allowing users to pipe code directly into the editor for formatting and linting. This tool provides a convenient way to integrate PHP CS Fixer with such editors.

## Usage

Since this is a wrapper to php-cs-fixer fix command then you can use any flag
defined in this tool such as `--config` and `--cache-file`.

### Zed

With the Zed editor, you can use this tool by adding the [configuration](zed-configuration.json) to your settings json.
