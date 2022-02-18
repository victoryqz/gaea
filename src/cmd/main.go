package main

import (
	"gaea.olympus.io/src/controllers"
	"gaea.olympus.io/src/core"
	"gaea.olympus.io/src/gaea"
	"gaea.olympus.io/src/middlewares"
)

func main() {
	core.GetK8sInformers().
		Mount(
			core.GetDeploymentHandler(),
			core.GetNamespaceHandler(),
		).
		Launch()

	gaea.Ignite().
		Attach(middlewares.NewRBACMid()).
		Mount(
			"/v01",
			controllers.GetNamespaceCtl(),
			controllers.GetWsCtl(),
		).
		Mount(
			"/v02",
			controllers.NewIndex(),
		).
		Launch()
}
