# FrankenPHP Queue

A [FrankenPHP](https://frankenphp.dev) extension that allows you to send messages in queues and handle them asynchronously.

It can be used as a lightweight, in-process replacement for queues systems like RabbitMQ, Beanstalkd and Redis
and can be used with [Symfony Messenger](https://symfony.com/doc/current/messenger.html) and [Laravel Queues](https://laravel.com/docs/12.x/queues).

> [!WARNING]
>
> This extension is highly experimental and not recommended for production use.
> The public API may change at any time without notice.

## Installation

First, if not already done, follow [the instructions to install a ZTS version of libphp and `xcaddy`](https://frankenphp.dev/docs/compile/#install-php).
Then, use [`xcaddy`](https://github.com/caddyserver/xcaddy) to build FrankenPHP with the `frankenphp-etcd` module:

```console
CGO_ENABLED=1 \
CGO_CFLAGS=$(php-config --includes) \
CGO_LDFLAGS="$(php-config --ldflags) $(php-config --libs)" \
xcaddy build \
    --output frankenphp \
    --with github.com/dunglas/frankenphp-queue \
    --with github.com/dunglas/frankenphp/caddy \
    --with github.com/dunglas/mercure/caddy \
    --with github.com/dunglas/vulcain/caddy
    # Add extra Caddy modules and FrankenPHP extensions here
```

That's all! Your custom FrankenPHP build contains the `frankenphp-queue` extension.

## Usage


### Register The Queue

Register the queue in your `Caddyfile`:

```caddyfile
{
	frankenphp
	frankenphp_queue {
 		# All directives are optional
 		worker queue-worker.php
 		name m#Queue
 		size 10000
 		min_threads 32 # defaults to the number of CPUs of the machine
    }
}

localhost {
    root public/
	php_server
}
```

### Write The Worker Script

```php
<?php

// queue-worker.php

// Handler outside the loop for better performance (doing less work)
$handler = static function (mixed $data)  {
    // Your logic here	
};

$maxRequests = (int)($_SERVER['MAX_REQUESTS'] ?? 0);
for ($nbRequests = 0; !$maxRequests || $nbRequests < $maxRequests; ++$nbRequests) {
    $keepRunning = \frankenphp_handle_request($handler);

    // Call the garbage collector to reduce the chances of it being triggered in the middle of the handling of a request
    gc_collect_cycles();

    if (!$keepRunning) {
      break;
    }
}
```

### Dispatch Messages

```php
<?php

// public/index.php

frankenphp_queue('Hello, Kévin!');

echo 'Data dispatched to an async worker.';
```

## Symfony Messenger Transport

Transport for Symfony Messenger is provided in tests,
it will be contributed to Symfony when the extension will be more mature.

* [`Caddyfile`](testdata/symfony-messenger/Caddyfile)
* [`queue-worker.php`](testdata/symfony-messenger/queue-worker.php)
* [`FrankenPHPQueueTransportFactory`](testdata/symfony-messenger/src/Transport/FrankenPHPQueueTransportFactory.php)
* [`FrankenPHPQueueTransport`](testdata/symfony-messenger/src/Transport/FrankenPHPQueueTransport.php)

## Credits

Created by [Kévin Dunglas](https://dunglas.dev) and sponsored by [Les-Tilleuls.coop](https://les-tilleuls.coop).
