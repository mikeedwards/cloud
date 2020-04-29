// Code generated by goa v3.1.2, DO NOT EDIT.
//
// HTTP request path constructors for the project service.
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	"fmt"
)

// UpdateProjectPath returns the URL path to the project service update HTTP endpoint.
func UpdateProjectPath(id int64) string {
	return fmt.Sprintf("/projects/%v/update", id)
}

// InvitesProjectPath returns the URL path to the project service invites HTTP endpoint.
func InvitesProjectPath() string {
	return "/projects/invites/pending"
}

// LookupInviteProjectPath returns the URL path to the project service lookup invite HTTP endpoint.
func LookupInviteProjectPath(token string) string {
	return fmt.Sprintf("/projects/invites/%v", token)
}

// AcceptInviteProjectPath returns the URL path to the project service accept invite HTTP endpoint.
func AcceptInviteProjectPath(id int64) string {
	return fmt.Sprintf("/projects/invites/%v/accept", id)
}

// RejectInviteProjectPath returns the URL path to the project service reject invite HTTP endpoint.
func RejectInviteProjectPath(id int64) string {
	return fmt.Sprintf("/projects/invites/%v/reject", id)
}