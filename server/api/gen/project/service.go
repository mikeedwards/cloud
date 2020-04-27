// Code generated by goa v3.1.2, DO NOT EDIT.
//
// project service
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package project

import (
	"context"

	projectviews "github.com/fieldkit/cloud/server/api/gen/project/views"
	"goa.design/goa/v3/security"
)

// Service is the project service interface.
type Service interface {
	// Update implements update.
	Update(context.Context, *UpdatePayload) (err error)
	// Invites implements invites.
	Invites(context.Context, *InvitesPayload) (res *PendingInvites, err error)
	// LookupInvite implements lookup invite.
	LookupInvite(context.Context, *LookupInvitePayload) (res *PendingInvites, err error)
	// AcceptInvite implements accept invite.
	AcceptInvite(context.Context, *AcceptInvitePayload) (err error)
	// RejectInvite implements reject invite.
	RejectInvite(context.Context, *RejectInvitePayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "project"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"update", "invites", "lookup invite", "accept invite", "reject invite"}

// UpdatePayload is the payload type of the project service update method.
type UpdatePayload struct {
	Auth string
	ID   int64
	Body string
}

// InvitesPayload is the payload type of the project service invites method.
type InvitesPayload struct {
	Auth string
}

// PendingInvites is the result type of the project service invites method.
type PendingInvites struct {
	Pending []*PendingInvite
}

// LookupInvitePayload is the payload type of the project service lookup invite
// method.
type LookupInvitePayload struct {
	Auth  string
	Token string
}

// AcceptInvitePayload is the payload type of the project service accept invite
// method.
type AcceptInvitePayload struct {
	Auth  string
	ID    int64
	Token *string
}

// RejectInvitePayload is the payload type of the project service reject invite
// method.
type RejectInvitePayload struct {
	Auth  string
	ID    int64
	Token *string
}

type PendingInvite struct {
	ID      int64
	Project *ProjectSummary
	Time    int64
}

type ProjectSummary struct {
	ID   int64
	Name string
}

// credentials are invalid
type Unauthorized string

// not found
type NotFound string

// Error returns an error description.
func (e Unauthorized) Error() string {
	return "credentials are invalid"
}

// ErrorName returns "unauthorized".
func (e Unauthorized) ErrorName() string {
	return "unauthorized"
}

// Error returns an error description.
func (e NotFound) Error() string {
	return "not found"
}

// ErrorName returns "not-found".
func (e NotFound) ErrorName() string {
	return "not-found"
}

// NewPendingInvites initializes result type PendingInvites from viewed result
// type PendingInvites.
func NewPendingInvites(vres *projectviews.PendingInvites) *PendingInvites {
	return newPendingInvites(vres.Projected)
}

// NewViewedPendingInvites initializes viewed result type PendingInvites from
// result type PendingInvites using the given view.
func NewViewedPendingInvites(res *PendingInvites, view string) *projectviews.PendingInvites {
	p := newPendingInvitesView(res)
	return &projectviews.PendingInvites{Projected: p, View: "default"}
}

// newPendingInvites converts projected type PendingInvites to service type
// PendingInvites.
func newPendingInvites(vres *projectviews.PendingInvitesView) *PendingInvites {
	res := &PendingInvites{}
	if vres.Pending != nil {
		res.Pending = make([]*PendingInvite, len(vres.Pending))
		for i, val := range vres.Pending {
			res.Pending[i] = transformProjectviewsPendingInviteViewToPendingInvite(val)
		}
	}
	return res
}

// newPendingInvitesView projects result type PendingInvites to projected type
// PendingInvitesView using the "default" view.
func newPendingInvitesView(res *PendingInvites) *projectviews.PendingInvitesView {
	vres := &projectviews.PendingInvitesView{}
	if res.Pending != nil {
		vres.Pending = make([]*projectviews.PendingInviteView, len(res.Pending))
		for i, val := range res.Pending {
			vres.Pending[i] = transformPendingInviteToProjectviewsPendingInviteView(val)
		}
	}
	return vres
}

// transformProjectviewsPendingInviteViewToPendingInvite builds a value of type
// *PendingInvite from a value of type *projectviews.PendingInviteView.
func transformProjectviewsPendingInviteViewToPendingInvite(v *projectviews.PendingInviteView) *PendingInvite {
	if v == nil {
		return nil
	}
	res := &PendingInvite{
		ID:   *v.ID,
		Time: *v.Time,
	}
	if v.Project != nil {
		res.Project = transformProjectviewsProjectSummaryViewToProjectSummary(v.Project)
	}

	return res
}

// transformProjectviewsProjectSummaryViewToProjectSummary builds a value of
// type *ProjectSummary from a value of type *projectviews.ProjectSummaryView.
func transformProjectviewsProjectSummaryViewToProjectSummary(v *projectviews.ProjectSummaryView) *ProjectSummary {
	res := &ProjectSummary{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// transformPendingInviteToProjectviewsPendingInviteView builds a value of type
// *projectviews.PendingInviteView from a value of type *PendingInvite.
func transformPendingInviteToProjectviewsPendingInviteView(v *PendingInvite) *projectviews.PendingInviteView {
	res := &projectviews.PendingInviteView{
		ID:   &v.ID,
		Time: &v.Time,
	}
	if v.Project != nil {
		res.Project = transformProjectSummaryToProjectviewsProjectSummaryView(v.Project)
	}

	return res
}

// transformProjectSummaryToProjectviewsProjectSummaryView builds a value of
// type *projectviews.ProjectSummaryView from a value of type *ProjectSummary.
func transformProjectSummaryToProjectviewsProjectSummaryView(v *ProjectSummary) *projectviews.ProjectSummaryView {
	res := &projectviews.ProjectSummaryView{
		ID:   &v.ID,
		Name: &v.Name,
	}

	return res
}
