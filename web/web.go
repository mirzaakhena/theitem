package web

import "embed"

//go:embed dist/*
var StaticFiles embed.FS

// contentStatic, _ := fs.Sub(fs.FS(web.StaticFiles), "dist")
// r.Router.StaticFS("/web", http.FS(contentStatic))

// https://stackoverflow.com/questions/62293398/cant-serve-vue-js-spa-app-using-the-noroute-function
