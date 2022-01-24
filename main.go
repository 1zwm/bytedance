package main

import (
	"bytedance/model"
	"bytedance/routes"
	"bytedance/utils"
)

func main() {
	g := routes.InitRouters()

	routes.RegisterRouter(g)
	model.InitDb()
	g.Run(utils.HttpPort)

}
