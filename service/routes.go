package service

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var RelationsRoutes = Routes{
	Route{
		Name:        "NewObject",
		Method:      "POST,OPTIONS",
		Pattern:     "/object/new",
		HandlerFunc: NewObject,
	},

	Route{
		Name:        "DeleteObject",
		Method:      "POST,OPTIONS",
		Pattern:     "/object/delete",
		HandlerFunc: DeleteObject,
	},

	Route{
		Name:        "NewRelation",
		Method:      "POST,OPTIONS",
		Pattern:     "/relation/new",
		HandlerFunc: NewRelation,
	},

	Route{
		Name:        "DeleteRelation",
		Method:      "POST,OPTIONS",
		Pattern:     "/relation/delete",
		HandlerFunc: DeleteRelation,
	},

	Route{
		Name:        "CheckRelationExistence",
		Method:      "GET,OPTIONS",
		Pattern:     "/relation/existence",
		HandlerFunc: CheckRelationExistence,
	},

	Route{
		Name:        "CheckRelationsExistence",
		Method:      "GET,OPTIONS",
		Pattern:     "/relations/existence",
		HandlerFunc: CheckRelationsExistence,
	},
}
