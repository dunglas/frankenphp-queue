<?php

declare(strict_types=1);

namespace App\Message;

final readonly class Notif
{
    public function __construct(
        public string $content,
    ) {}
}
