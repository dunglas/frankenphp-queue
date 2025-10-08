package queue

import (
	"runtime"
	"strconv"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/dunglas/frankenphp"
	"go.uber.org/zap"
)

func init() {
	caddy.RegisterModule(Queue{})
	httpcaddyfile.RegisterGlobalOption("frankenphp_queue", parseGlobalOption)
}

type Queue struct {
	Size       int    `json:"size,omitempty"`
	MinThreads int    `json:"min_threads,omitempty"`
	Name       string `json:"name,omitempty"`
	Worker     string `json:"worker,omitempty"`

	ctx    caddy.Context
	logger *zap.Logger
}

// CaddyModule returns the Caddy module information.
func (Queue) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "frankenphp_queue",
		New: func() caddy.Module { return new(Queue) },
	}
}

func (g *Queue) Provision(ctx caddy.Context) error {
	g.logger = ctx.Logger()
	g.ctx = ctx

	if g.Size <= 0 {
		g.Size = 10_000
	}

	if g.MinThreads <= 0 {
		g.MinThreads = runtime.NumCPU()
	}

	if g.Worker == "" {
		g.Worker = "queue-worker.php"
	}

	w.requestChan = make(chan *frankenphp.WorkerRequest, g.Size)
	w.minThread = g.MinThreads
	w.name = g.Name
	w.filename = g.Worker
	w.logger = g.logger

	frankenphp.RegisterWorker(w)

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
			case "min_threads":
				if !d.NextArg() {
					return d.ArgErr()
				}

				t, err := strconv.Atoi(d.Val())
				if err != nil {
					return d.Errf("failed to parse min_threads: %v", err)
				}
				g.MinThreads = t
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
