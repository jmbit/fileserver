package httpd

import (
	"io/fs"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/jmbit/fileserver/internal/middlewares"
	"github.com/jmbit/fileserver/frontend"
	"github.com/spf13/viper"
)

func RegisterRoutes() *http.ServeMux {
  mux := http.NewServeMux()
  
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!"))
  })
        //Check if running in dev mode, if yes, proxy requests not going to the API endpoints
        //to the vite dev server instead. In prod will also enable caching for static files
        if viper.GetBool("debug.devmode") {
                targetUrl, err := url.Parse("http://localhost:5173/")
                if err != nil {
                        panic(err)
                }
                reverseProxy := httputil.NewSingleHostReverseProxy(targetUrl)
                mux.Handle("/", reverseProxy)
        } else {
                fsroot, err := fs.Sub(frontend.FrontendFS, "dist")
                if err != nil {
                        panic(err)
                }
                fileServer := http.FileServer(http.FS(fsroot))
                mux.Handle("/", middlewares.AssetCaching(fileServer))
        }
  
  return mux

}
