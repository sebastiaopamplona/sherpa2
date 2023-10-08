package handler

import "github.com/sebastiaopamplona/sherpa2/api/openapi"

type Handler struct {
	// TODO: add config
	// TODO: add db client
	Routes openapi.ApiHandleFunctions
}

func New() *Handler {
	h := &Handler{}

	routes := openapi.ApiHandleFunctions{
		ProjectAPI: openapi.ProjectAPI{
			CreateProject:   h.CreateProject,
			DeleteProject:   h.DeleteProject,
			GetProject:      h.GetProject,
			GetProjectUsers: h.GetProjectUsers,
			ListProjects:    h.ListProjects,
			UpdateProject:   h.UpdateProject,
		},
		SprintAPI: openapi.SprintAPI{
			CreateSprint:            h.CreateSprint,
			DeleteSprint:            h.DeleteSprint,
			GetSprint:               h.GetSprint,
			GetSprintStoryBreakdown: h.GetSprintStoryBreakdown,
			GetSprintUserBreakdown:  h.GetSprintUserBreakdown,
			ListSprints:             h.ListSprints,
			UpdateSprint:            h.UpdateSprint,
		},
		StoryAPI: openapi.StoryAPI{
			CreateStory:        h.CreateStory,
			DeleteStory:        h.DeleteStory,
			GetStory:           h.GetStory,
			GetStoryTimekeeper: h.GetStoryTimekeeper,
			ListStories:        h.ListStories,
			UpdateStory:        h.UpdateStory,
		},
	}

	h.Routes = routes

	return h
}
