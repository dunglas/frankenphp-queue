<?php

declare(strict_types=1);

namespace App\Transport;

use Symfony\Component\Messenger\Envelope;
use Symfony\Component\Messenger\Exception\LogicException;
use Symfony\Component\Messenger\Transport\Serialization\PhpSerializer;
use Symfony\Component\Messenger\Transport\Serialization\SerializerInterface;
use Symfony\Component\Messenger\Transport\TransportInterface;

/**
 * @author Kévin Dunglas <kevin@dunglas.dev>
 */
class FrankenPHPQueueTransport implements TransportInterface
{
    private SerializerInterface $serializer;

    public function __construct(
        ?SerializerInterface $serializer = null,
    ) {
        $this->serializer = $serializer ?? new PhpSerializer();
    }

    public function get(): iterable
    {
        $envelope = null;
        \frankenphp_handle_request(function (string $message) use (&$envelope) {
            $envelope = $this->serializer->decode([
                'body' => $message,
            ]);
        });

        return $envelope;
    }

    public function ack(Envelope $envelope): void
    {
        throw new LogicException('Not implemented');
    }

    public function reject(Envelope $envelope): void
    {
        throw new LogicException('Not implemented');
    }

    public function send(Envelope $envelope): Envelope
    {
        \frankenphp_queue($this->serializer->encode($envelope));

        return $envelope;
    }
}