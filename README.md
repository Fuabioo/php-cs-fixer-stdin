# php-cs-fixer-stdin

[PHP CS Fixer](https://github.com/PHP-CS-Fixer/PHP-CS-Fixer) wrapper for stdin text editor integration

## Context

Some code editors support stdin integration, allowing users to pipe code directly into the editor for formatting and linting. This tool provides a convenient way to integrate PHP CS Fixer with such editors.

## Installation

```sh
wget https://github.com/Fuabioo/php-cs-fixer-stdin/releases/latest/download/php-cs-fixer-stdin_$(uname -s)_$(uname -m).tar.gz
tar -xzf php-cs-fixer-stdin_$(uname -s)_$(uname -m).tar.gz
chmod +x php-cs-fixer-stdin
```

## Usage

Since this is a wrapper to php-cs-fixer fix command then you can use any flag
defined in this tool such as `--config` and `--cache-file`.

### Zed

With the [Zed editor](https://zed.dev/), you can use this tool by adding the [configuration](zed-configuration.json) to your settings json.

How [Zed formatting](https://zed.dev/docs/configuring-zed#formatter) works for anything other than a language server formatter is that you set up an external CLI tool, to which it will pipe the code. Currently there is no support for providing something like a filepath to the external tool or the like.

> For more information about PHP language support, see [PHP - Zed](https://zed.dev/docs/languages/php).

### Developing

If you have the go/bin directory in your PATH, you can simply run `go install` to install the tool from the source code in your local machine.
Once that is done, just pipe any php code into the tool or
(if you already have one) use your configured text editor.

Using the following configuration in an `example-config.php` file:

```php
<?php

return (new PhpCsFixer\Config())
    ->setRiskyAllowed(true)
    ->setRules([
        'psr_autoloading' => false,
        '@PSR12' => true,
        'trim_array_spaces' => true,
    ]);
```

Execute the tool with the following command:

```sh
echo '<?php
$sample = array( 'a', 'b' );' | php-cs-fixer-stdin --config=example-config.php --using-cache=no
```

This should throw the following output (the editor will know what to do with it):

```php
<?php

$sample = array(a, b);
```
