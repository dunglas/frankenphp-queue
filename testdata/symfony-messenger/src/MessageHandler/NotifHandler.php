<?php

declare(strict_types=1);

namespace App\MessageHandler;

use App\Message\Notif;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;

#[AsMessageHandler]
final class NotifHandler
{
    public function __invoke(Notif $notif): void
    {
        error_log($notif->content, 4);
    }
}
