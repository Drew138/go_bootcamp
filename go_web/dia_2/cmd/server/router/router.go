package router

import (
	"dia_2/cmd/server/handlers"
	"dia_2/cmd/server/middleware"
	"dia_2/internal/product"
	"dia_2/pkg/store"
	"os"

	"github.com/gin-gonic/gin"
)

type Router struct {
	en *gin.Engine
}

func NewRouter(e *gin.Engine) Router {
	return Router{e}
}

func (r *Router) SetRoutes() {
	pr := r.en.Group("/product")
	var PRODUCT_STORE_FILE = os.Getenv("PRODUCT_STORE_FILE")
	pst := store.NewProductStore(PRODUCT_STORE_FILE)
	rp := product.NewRepository(pst)
	sv := product.NewService(rp)
	phm := handlers.NewProductHandlerManager(sv)

	pr.Use(middleware.IsAuthenticated())
	pr.GET("/ping", phm.Ping())
	pr.GET("/:id", phm.GetById())
	pr.GET("", phm.Get())
	pr.POST("", phm.Create())
	pr.DELETE("/:id", phm.Delete())
	pr.PATCH("/:id", phm.Patch())
	pr.PUT("/:id", phm.Put())
}
