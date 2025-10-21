package templates

import "github.com/gin-contrib/multitemplate"

func SetupTemplates() multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()

	base := "ui/template/base.html"

	// home and error
	renderer.AddFromFiles("home", base, "ui/template/home.html")
	renderer.AddFromFiles("error", base, "ui/template/error.html")

	// auth
	authPath := "ui/template/auth/"

	renderer.AddFromFiles("login", base, authPath+"login.html")
	renderer.AddFromFiles("register", base, authPath+"register.html")

	return renderer

}