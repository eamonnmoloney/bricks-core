package controllers

import "github.com/gin-gonic/gin"
import "net/http"

func ReadComponents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// CreateComponent The component will
// add the component to the application, create the git repo for the module. This can be a nextjs app or an API
// All new components are created in the sandbox first and application promotion workflows bring them to high environments
func CreateComponent(c *gin.Context) {
	// Required: name, type
	// The application ID will be taken from the path
	// /applications/application_id/components
	// I can do this when I have the ABAC
	// subject: engineer | qa
	// environment: sandbox | prod
	// action: create | read | update | delete
	// resource: *
	// resource_type: component | application
	// team: green | red | etc.
	// sensitivity: if the requirement is to make a component that is sensitive, the user must have the required attribute for this. E.g. HIPPA, GDPR, PII, Custom.
	// So user will need to have these attributes.
	// This can be in a flat file as this should be configured using terraform
	c.JSON(http.StatusCreated, gin.H{
		"message": "pong",
	})
}
