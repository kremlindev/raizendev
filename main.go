package main
import(
  "fmt"
  "net/http"
  "github.com/labstack/echo/v4"
)
func main() {
  fmt.Println("-------------------------------")
  fmt.Println("🎯 Proje Başarıyla Başlatıldı")
  fmt.Println("🎯 Herhangi bir sorun bulunamadı")
  fmt.Println("-------------------------------")
  e := echo.New()
  e.GET("/", func(c echo.Context) error {
    return c.File("./src/main.html")
  })
  e.GET("/css", func(c echo.Context) error {
    return c.File("main.css")
  })
  e.GET("/redirect/:name", func(c echo.Context) error {
    if c.Param("name") == "github" {
      return c.Redirect(http.StatusMovedPermanently, "https://github.com/raizendev")
    }
    if c.Param("name") == "spotify" {
      return c.Redirect(http.StatusMovedPermanently, "https://open.spotify.com/user/cg68ppjbrdbneurrkftnb65ye?si=gbT5afAaRFeC2IZmTR69LA&utm_source=copy-link&dl_branch=1")
    }
    if c.Param("name") == "discord" {
      return c.Redirect(http.StatusMovedPermanently, "https://discord.com/users/807863820977831946")
    }
    return c.Redirect(http.StatusMovedPermanently, "/")
  })
  e.GET("/api", func(c echo.Context) error {
    if c.QueryParam("api") == "api" {
      c.JSON('{"status": "200"}') 
    }
     if error != nil {
     c.JSON('{"status": "400"}')  
    }
    c.JSON('{"status": "204"}')
  })
  e.Logger.Fatal(e.Start(":3000"))
}
