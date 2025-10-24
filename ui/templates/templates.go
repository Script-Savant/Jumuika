package templates

import "github.com/gin-contrib/multitemplate"

func SetupTemplates() multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()

	base := "ui/templates/base.html"

	// home and error
	renderer.AddFromFiles("home", base, "ui/templates/home.html")
	renderer.AddFromFiles("error", base, "ui/templates/error.html")

	// auth
	authPath := "ui/templates/auth/"
	renderer.AddFromFiles("login", base, authPath+"login.html")
	renderer.AddFromFiles("register", base, authPath+"register.html")

	// profile
	profilePath := "ui/templates/profile"
	renderer.AddFromFiles("profile-view", base, profilePath+"view.html")

	return renderer

}
