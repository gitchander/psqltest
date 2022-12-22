package main

import (
	"html/template"
	"log"
	"net/http"
)

type rootHandler struct {
	hc *HandleCore
}

var _ http.Handler = &rootHandler{}

func (h *rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	pageMenu := []PageMenuItem{
		{
			Name: "Sign Up",
			Href: "/users/sign_up",
		},
		{
			Name: "Sign In",
			Href: "/users/sign_in",
		},
	}

	// pageMenu := []PageMenuItem{
	// 	{
	// 		Name: "Sign Down",
	// 		Href: "/sign_up",
	// 	},
	// 	{
	// 		Name: "Sign Out",
	// 		Href: "/sign_in",
	// 	},
	// }

	files := []string{
		"assets/templates/base.ate",
		"assets/templates/page_menu.ate",
		"assets/templates/content_header.ate",
		"assets/templates/breadcrumbs.ate",
		"assets/templates/root.ate",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		errLog(err)
		return
	}

	params := map[string]interface{}{
		"hasTitle": false,
		"styles": []string{
			"/assets/styles/index.css",
			"/assets/styles/page_menu.css",
			"/assets/styles/breadcrumbs.css",
		},
		"breadcrumbs": []Breadcrumb{
			{
				Name: "Home",
				Href: "/",
			},
		},
		"contentHeader": ContentHeader{
			Title: "Home",
		},
		"pageMenu": pageMenu,
	}

	err = t.Execute(w, params)
	if handleError(err) {
		return
	}
}

// ------------------------------------------------------------------------------
// sign up - зарегистрироваться
type signUpHandler struct {
	hc *HandleCore
}

var _ http.Handler = &signUpHandler{}

func (h *signUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Println("Method:", r.Method)
	if r.Method == "POST" {
		var (
			email           = r.FormValue("email")
			username        = r.FormValue("username")
			password        = r.FormValue("password")
			passwordConfirm = r.FormValue("password_confirm")
		)
		log.Printf("email:%q, username:%q, password:%q, passwordConfirm:%q\n", email, username, password, passwordConfirm)
	}

	pageMenu := []PageMenuItem{
		// {
		// 	Name: "Sign Up",
		// 	Href: "/users/sign_up",
		// },
		{
			Name: "Sign In",
			Href: "/users/sign_in",
		},
	}

	files := []string{
		"assets/templates/base.ate",
		"assets/templates/page_menu.ate",
		"assets/templates/content_header.ate",
		"assets/templates/breadcrumbs.ate",
		"assets/templates/sign_up.ate",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		errLog(err)
		return
	}

	params := map[string]interface{}{
		"hasTitle": false,
		"styles": []string{
			// "/assets/styles/index.css",
			// "/assets/styles/page_menu.css",
			// "/assets/styles/breadcrumbs.css",
			"/assets/styles/sign_up.css",
		},
		"breadcrumbs": []Breadcrumb{
			{
				Name: "Home",
				Href: "/",
			},
		},
		"contentHeader": ContentHeader{
			Title: "Home",
		},
		"pageMenu": pageMenu,
	}

	err = t.Execute(w, params)
	if handleError(err) {
		return
	}
}

// ------------------------------------------------------------------------------
// sign down - разрегистрироваться, отписаться
type signDownHandler struct {
	hc *HandleCore
}

var _ http.Handler = &signDownHandler{}

func (h *signDownHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

//------------------------------------------------------------------------------
// https://www.codewars.com/users/sign_in
// https://www.w3schools.com/howto/howto_css_login_form.asp

// sign in - войти
type signInHandler struct {
	hc *HandleCore
}

var _ http.Handler = &signInHandler{}

func (h *signInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

// ------------------------------------------------------------------------------
// sign out - выход
type signOutHandler struct {
	hc *HandleCore
}

var _ http.Handler = &signOutHandler{}

func (h *signOutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

// ------------------------------------------------------------------------------
type workHandler struct {
	hc *HandleCore
}

var _ http.Handler = &workHandler{}

func (h *workHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
