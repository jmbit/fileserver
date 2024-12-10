package httpd

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jmbit/fileserver/internal/middlewares"
	"github.com/spf13/viper"
)

func NewServer() *http.Server {
        port := viper.GetInt("web.port")
        host := viper.GetString("web.host")
        middlewareStack := middlewares.CreateStack()
        if viper.GetBool("debug.log.http") {
                log.Println("Enabling HTTP logging")
                middlewareStack = middlewares.CreateStack(
                        middlewares.GorillaLogging,
                )
        } else {
                middlewareStack = middlewares.CreateStack(
                )
        }

        // Declare Server config
        server := &http.Server{
                Addr:         fmt.Sprintf("%s:%d", host, port),
                Handler:      middlewareStack(RegisterRoutes()),
                IdleTimeout:  time.Minute,
                ReadTimeout:  10 * time.Second,
                WriteTimeout: 30 * time.Second,
        }

        return server
}
