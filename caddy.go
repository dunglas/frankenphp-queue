package queue

import (
	"log/slog"
	"strconv"
	"sync"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/dunglas/frankenphp"
	frankenphpCaddy "github.com/dunglas/frankenphp/caddy"
)

var (
	worker   frankenphp.Workers
	logger   *slog.Logger
	workerMu sync.Mutex
)

func init() {
	caddy.RegisterModule(Queue{})
	httpcaddyfile.RegisterGlobalOption("frankenphp_queue", parseGlobalOption)
}

type Queue struct {
	Size       int    `json:"size,omitempty"`
	NumThreads int    `json:"numthreads,omitempty"`
	Name       string `json:"name,omitempty"`
	Worker     string `json:"worker,omitempty"`
}

// CaddyModule returns the Caddy module information.
func (Queue) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "frankenphp_queue",
		New: func() caddy.Module { return new(Queue) },
	}
}

func (g *Queue) Provision(ctx caddy.Context) error {
	if g.Size <= 0 {
		g.Size = 10_000
	}

	if g.Name == "" {
		g.Name = "m#Queue"
	}

	if g.Worker == "" {
		g.Worker = "queue-worker.php"
	}

	workerMu.Lock()
	worker = frankenphpCaddy.RegisterWorkers(g.Name, g.Worker, g.NumThreads)
	logger = ctx.Slogger()
	workerMu.Unlock()

	return nil
}

func (g *Queue) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		for d.NextBlock(0) {
			// when adding a new directive, also update the allowedDirectives error message
			switch d.Val() {
			case "worker":
				if !d.NextArg() {
					return d.ArgErr()
				}

				g.Worker = d.Val()
			case "name":
				if !d.NextArg() {
					return d.ArgErr()
				}

				g.Worker = d.Val()
			case "size":
				if !d.NextArg() {
					return d.ArgErr()
				}

				s, err := strconv.Atoi(d.Val())
				if err != nil {
					return d.Errf("failed to parse size: %v", err)
				}
				g.Size = s
			case "num":
				if !d.NextArg() {
					return d.ArgErr()
				}

				t, err := strconv.Atoi(d.Val())
				if err != nil {
					return d.Errf("failed to parse min_threads: %v", err)
				}
				g.NumThreads = t
			default:
				return d.Errf(`unrecognized subdirective "%s"`, d.Val())
			}
		}
	}

	return nil
}

func parseGlobalOption(d *caddyfile.Dispenser, _ any) (any, error) {
	app := &Queue{}
	if err := app.UnmarshalCaddyfile(d); err != nil {
		return nil, err
	}

	// tell Caddyfile adapter that this is the JSON for an app
	return httpcaddyfile.App{
		Name:  "frankenphp_queue",
		Value: caddyconfig.JSON(app, nil),
	}, nil
}

// Interface guards
var (
	_ caddy.Module = (*Queue)(nil)
)
