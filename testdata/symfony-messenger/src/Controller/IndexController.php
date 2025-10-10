<?php

declare(strict_types=1);

namespace App\Controller;

use App\Message\Notif;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Messenger\MessageBusInterface;
use Symfony\Component\Routing\Attribute\Route;

final class IndexController extends AbstractController
{
    #[Route('/')]
    public function index(MessageBusInterface $bus): Response {
        $bus->dispatch(new Notif('Yo!'));

        return new Response('Hello, FrankenPHP Queue!');
    }
}
