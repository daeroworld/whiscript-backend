package configuration

import (
	"fmt"
	"net/http"
	voiceReader "voice/reader/configuration"
	voiceWriter "voice/writer/configuration"

	sharedConfiguration "github.com/daeroworld/shared/configuration"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Container struct {
	Router      *gin.Engine
	Variable    *Variable
	voiceWriter sharedConfiguration.IContainer
	voiceReader sharedConfiguration.IContainer
}

func (c *Container) GetHttpHandler() http.Handler {
	return c.Router.Handler()
}

func (c *Container) InitVariable() error {
	c.Variable = NewVariable()
	return nil
}

func (c *Container) SetRouter(router any) {
	c.Router = router.(*gin.Engine)

	front := fmt.Sprintf("http://%s:%d", c.Variable.Frontend.Ip, c.Variable.Frontend.Port)
	c.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{front}, // Allow frontend domain
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
}

func (c *Container) DefineDatabase(databaseWrappers ...any) error {

	return nil
}

func (c *Container) DefineRoute(router interface{}) error {
	//c.voiceWriter.DefineRoute(router)
	//c.voiceReader.DefineRoute(router)
	return nil
}

func (c *Container) DefineGrpc() error {
	panic("not implemented")
}

func (c *Container) InitDependency(db interface{}) error {
	c.voiceWriter = voiceWriter.NewMonoContainer()
	c.voiceReader = voiceReader.NewMonoContainer()
	return nil
}

func NewContainer() *Container {
	ctnr := &Container{}
	ctnr.InitVariable()
	ctnr.DefineDatabase(nil)
	ctnr.InitDependency(nil)
	router := gin.Default()
	ctnr.SetRouter(router)
	ctnr.DefineRoute(router)
	return ctnr
}
