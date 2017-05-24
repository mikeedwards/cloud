// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "fieldkit": Application User Types
//
// Command:
// $ main

package app

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// addAdministratorPayload user type.
type addAdministratorPayload struct {
	UserID *int `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
}

// Validate validates the addAdministratorPayload type instance.
func (ut *addAdministratorPayload) Validate() (err error) {
	if ut.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "user_id"))
	}
	return
}

// Publicize creates AddAdministratorPayload from addAdministratorPayload
func (ut *addAdministratorPayload) Publicize() *AddAdministratorPayload {
	var pub AddAdministratorPayload
	if ut.UserID != nil {
		pub.UserID = *ut.UserID
	}
	return &pub
}

// AddAdministratorPayload user type.
type AddAdministratorPayload struct {
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
}

// addExpeditionPayload user type.
type addExpeditionPayload struct {
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Slug        *string `form:"slug,omitempty" json:"slug,omitempty" xml:"slug,omitempty"`
}

// Validate validates the addExpeditionPayload type instance.
func (ut *addExpeditionPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Slug == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "slug"))
	}
	if ut.Description == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if ut.Slug != nil {
		if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, *ut.Slug); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.slug`, *ut.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
		}
	}
	if ut.Slug != nil {
		if utf8.RuneCountInString(*ut.Slug) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.slug`, *ut.Slug, utf8.RuneCountInString(*ut.Slug), 40, false))
		}
	}
	return
}

// Publicize creates AddExpeditionPayload from addExpeditionPayload
func (ut *addExpeditionPayload) Publicize() *AddExpeditionPayload {
	var pub AddExpeditionPayload
	if ut.Description != nil {
		pub.Description = *ut.Description
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Slug != nil {
		pub.Slug = *ut.Slug
	}
	return &pub
}

// AddExpeditionPayload user type.
type AddExpeditionPayload struct {
	Description string `form:"description" json:"description" xml:"description"`
	Name        string `form:"name" json:"name" xml:"name"`
	Slug        string `form:"slug" json:"slug" xml:"slug"`
}

// Validate validates the AddExpeditionPayload type instance.
func (ut *AddExpeditionPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Slug == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "slug"))
	}
	if ut.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, ut.Slug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.slug`, ut.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
	}
	if utf8.RuneCountInString(ut.Slug) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.slug`, ut.Slug, utf8.RuneCountInString(ut.Slug), 40, false))
	}
	return
}

// addFieldkitInputPayload user type.
type addFieldkitInputPayload struct {
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the addFieldkitInputPayload type instance.
func (ut *addFieldkitInputPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// Publicize creates AddFieldkitInputPayload from addFieldkitInputPayload
func (ut *addFieldkitInputPayload) Publicize() *AddFieldkitInputPayload {
	var pub AddFieldkitInputPayload
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	return &pub
}

// AddFieldkitInputPayload user type.
type AddFieldkitInputPayload struct {
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the AddFieldkitInputPayload type instance.
func (ut *AddFieldkitInputPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// addMemberPayload user type.
type addMemberPayload struct {
	Role   *string `form:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
	UserID *int    `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
}

// Validate validates the addMemberPayload type instance.
func (ut *addMemberPayload) Validate() (err error) {
	if ut.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "user_id"))
	}
	if ut.Role == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "role"))
	}
	return
}

// Publicize creates AddMemberPayload from addMemberPayload
func (ut *addMemberPayload) Publicize() *AddMemberPayload {
	var pub AddMemberPayload
	if ut.Role != nil {
		pub.Role = *ut.Role
	}
	if ut.UserID != nil {
		pub.UserID = *ut.UserID
	}
	return &pub
}

// AddMemberPayload user type.
type AddMemberPayload struct {
	Role   string `form:"role" json:"role" xml:"role"`
	UserID int    `form:"user_id" json:"user_id" xml:"user_id"`
}

// Validate validates the AddMemberPayload type instance.
func (ut *AddMemberPayload) Validate() (err error) {

	if ut.Role == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "role"))
	}
	return
}

// addProjectPayload user type.
type addProjectPayload struct {
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Slug        *string `form:"slug,omitempty" json:"slug,omitempty" xml:"slug,omitempty"`
}

// Validate validates the addProjectPayload type instance.
func (ut *addProjectPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Slug == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "slug"))
	}
	if ut.Description == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if ut.Slug != nil {
		if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, *ut.Slug); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.slug`, *ut.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
		}
	}
	if ut.Slug != nil {
		if utf8.RuneCountInString(*ut.Slug) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.slug`, *ut.Slug, utf8.RuneCountInString(*ut.Slug), 40, false))
		}
	}
	return
}

// Publicize creates AddProjectPayload from addProjectPayload
func (ut *addProjectPayload) Publicize() *AddProjectPayload {
	var pub AddProjectPayload
	if ut.Description != nil {
		pub.Description = *ut.Description
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Slug != nil {
		pub.Slug = *ut.Slug
	}
	return &pub
}

// AddProjectPayload user type.
type AddProjectPayload struct {
	Description string `form:"description" json:"description" xml:"description"`
	Name        string `form:"name" json:"name" xml:"name"`
	Slug        string `form:"slug" json:"slug" xml:"slug"`
}

// Validate validates the AddProjectPayload type instance.
func (ut *AddProjectPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Slug == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "slug"))
	}
	if ut.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, ut.Slug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.slug`, ut.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
	}
	if utf8.RuneCountInString(ut.Slug) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.slug`, ut.Slug, utf8.RuneCountInString(ut.Slug), 40, false))
	}
	return
}

// addTeamPayload user type.
type addTeamPayload struct {
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Slug        *string `form:"slug,omitempty" json:"slug,omitempty" xml:"slug,omitempty"`
}

// Validate validates the addTeamPayload type instance.
func (ut *addTeamPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Slug == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "slug"))
	}
	if ut.Description == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if ut.Name != nil {
		if ok := goa.ValidatePattern(`\S`, *ut.Name); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.name`, *ut.Name, `\S`))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) > 256 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 256, false))
		}
	}
	if ut.Slug != nil {
		if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, *ut.Slug); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.slug`, *ut.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
		}
	}
	if ut.Slug != nil {
		if utf8.RuneCountInString(*ut.Slug) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.slug`, *ut.Slug, utf8.RuneCountInString(*ut.Slug), 40, false))
		}
	}
	return
}

// Publicize creates AddTeamPayload from addTeamPayload
func (ut *addTeamPayload) Publicize() *AddTeamPayload {
	var pub AddTeamPayload
	if ut.Description != nil {
		pub.Description = *ut.Description
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Slug != nil {
		pub.Slug = *ut.Slug
	}
	return &pub
}

// AddTeamPayload user type.
type AddTeamPayload struct {
	Description string `form:"description" json:"description" xml:"description"`
	Name        string `form:"name" json:"name" xml:"name"`
	Slug        string `form:"slug" json:"slug" xml:"slug"`
}

// Validate validates the AddTeamPayload type instance.
func (ut *AddTeamPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Slug == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "slug"))
	}
	if ut.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if ok := goa.ValidatePattern(`\S`, ut.Name); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.name`, ut.Name, `\S`))
	}
	if utf8.RuneCountInString(ut.Name) > 256 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, ut.Name, utf8.RuneCountInString(ut.Name), 256, false))
	}
	if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, ut.Slug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.slug`, ut.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
	}
	if utf8.RuneCountInString(ut.Slug) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.slug`, ut.Slug, utf8.RuneCountInString(ut.Slug), 40, false))
	}
	return
}

// addTwitterAccountInputPayload user type.
type addTwitterAccountInputPayload struct {
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the addTwitterAccountInputPayload type instance.
func (ut *addTwitterAccountInputPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// Publicize creates AddTwitterAccountInputPayload from addTwitterAccountInputPayload
func (ut *addTwitterAccountInputPayload) Publicize() *AddTwitterAccountInputPayload {
	var pub AddTwitterAccountInputPayload
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	return &pub
}

// AddTwitterAccountInputPayload user type.
type AddTwitterAccountInputPayload struct {
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the AddTwitterAccountInputPayload type instance.
func (ut *AddTwitterAccountInputPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// addUserPayload user type.
type addUserPayload struct {
	Bio         *string `form:"bio,omitempty" json:"bio,omitempty" xml:"bio,omitempty"`
	Email       *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	InviteToken *string `form:"invite_token,omitempty" json:"invite_token,omitempty" xml:"invite_token,omitempty"`
	Name        *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Password    *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// Username
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
}

// Validate validates the addUserPayload type instance.
func (ut *addUserPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.Username == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "username"))
	}
	if ut.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "password"))
	}
	if ut.Bio == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "bio"))
	}
	if ut.InviteToken == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "invite_token"))
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Name != nil {
		if ok := goa.ValidatePattern(`\S`, *ut.Name); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.name`, *ut.Name, `\S`))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) > 256 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 256, false))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 10, true))
		}
	}
	if ut.Username != nil {
		if ok := goa.ValidatePattern(`^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`, *ut.Username); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.username`, *ut.Username, `^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`))
		}
	}
	if ut.Username != nil {
		if utf8.RuneCountInString(*ut.Username) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.username`, *ut.Username, utf8.RuneCountInString(*ut.Username), 40, false))
		}
	}
	return
}

// Publicize creates AddUserPayload from addUserPayload
func (ut *addUserPayload) Publicize() *AddUserPayload {
	var pub AddUserPayload
	if ut.Bio != nil {
		pub.Bio = *ut.Bio
	}
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.InviteToken != nil {
		pub.InviteToken = *ut.InviteToken
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Password != nil {
		pub.Password = *ut.Password
	}
	if ut.Username != nil {
		pub.Username = *ut.Username
	}
	return &pub
}

// AddUserPayload user type.
type AddUserPayload struct {
	Bio         string `form:"bio" json:"bio" xml:"bio"`
	Email       string `form:"email" json:"email" xml:"email"`
	InviteToken string `form:"invite_token" json:"invite_token" xml:"invite_token"`
	Name        string `form:"name" json:"name" xml:"name"`
	Password    string `form:"password" json:"password" xml:"password"`
	// Username
	Username string `form:"username" json:"username" xml:"username"`
}

// Validate validates the AddUserPayload type instance.
func (ut *AddUserPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "username"))
	}
	if ut.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "password"))
	}
	if ut.Bio == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "bio"))
	}
	if ut.InviteToken == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "invite_token"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, ut.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, ut.Email, goa.FormatEmail, err2))
	}
	if ok := goa.ValidatePattern(`\S`, ut.Name); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.name`, ut.Name, `\S`))
	}
	if utf8.RuneCountInString(ut.Name) > 256 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, ut.Name, utf8.RuneCountInString(ut.Name), 256, false))
	}
	if utf8.RuneCountInString(ut.Password) < 10 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, ut.Password, utf8.RuneCountInString(ut.Password), 10, true))
	}
	if ok := goa.ValidatePattern(`^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`, ut.Username); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.username`, ut.Username, `^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`))
	}
	if utf8.RuneCountInString(ut.Username) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.username`, ut.Username, utf8.RuneCountInString(ut.Username), 40, false))
	}
	return
}

// loginPayload user type.
type loginPayload struct {
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// Username
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
}

// Validate validates the loginPayload type instance.
func (ut *loginPayload) Validate() (err error) {
	if ut.Username == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "username"))
	}
	if ut.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "password"))
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 10, true))
		}
	}
	if ut.Username != nil {
		if ok := goa.ValidatePattern(`^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`, *ut.Username); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.username`, *ut.Username, `^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`))
		}
	}
	if ut.Username != nil {
		if utf8.RuneCountInString(*ut.Username) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.username`, *ut.Username, utf8.RuneCountInString(*ut.Username), 40, false))
		}
	}
	return
}

// Publicize creates LoginPayload from loginPayload
func (ut *loginPayload) Publicize() *LoginPayload {
	var pub LoginPayload
	if ut.Password != nil {
		pub.Password = *ut.Password
	}
	if ut.Username != nil {
		pub.Username = *ut.Username
	}
	return &pub
}

// LoginPayload user type.
type LoginPayload struct {
	Password string `form:"password" json:"password" xml:"password"`
	// Username
	Username string `form:"username" json:"username" xml:"username"`
}

// Validate validates the LoginPayload type instance.
func (ut *LoginPayload) Validate() (err error) {
	if ut.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "username"))
	}
	if ut.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "password"))
	}
	if utf8.RuneCountInString(ut.Password) < 10 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, ut.Password, utf8.RuneCountInString(ut.Password), 10, true))
	}
	if ok := goa.ValidatePattern(`^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`, ut.Username); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.username`, ut.Username, `^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`))
	}
	if utf8.RuneCountInString(ut.Username) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.username`, ut.Username, utf8.RuneCountInString(ut.Username), 40, false))
	}
	return
}

// updateFieldkitInputPayload user type.
type updateFieldkitInputPayload struct {
	Name   *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	TeamID *int    `form:"team_id,omitempty" json:"team_id,omitempty" xml:"team_id,omitempty"`
	UserID *int    `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
}

// Publicize creates UpdateFieldkitInputPayload from updateFieldkitInputPayload
func (ut *updateFieldkitInputPayload) Publicize() *UpdateFieldkitInputPayload {
	var pub UpdateFieldkitInputPayload
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	if ut.TeamID != nil {
		pub.TeamID = ut.TeamID
	}
	if ut.UserID != nil {
		pub.UserID = ut.UserID
	}
	return &pub
}

// UpdateFieldkitInputPayload user type.
type UpdateFieldkitInputPayload struct {
	Name   *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	TeamID *int    `form:"team_id,omitempty" json:"team_id,omitempty" xml:"team_id,omitempty"`
	UserID *int    `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
}

// updateInputPayload user type.
type updateInputPayload struct {
	Name   *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	TeamID *int    `form:"team_id,omitempty" json:"team_id,omitempty" xml:"team_id,omitempty"`
	UserID *int    `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
}

// Validate validates the updateInputPayload type instance.
func (ut *updateInputPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// Publicize creates UpdateInputPayload from updateInputPayload
func (ut *updateInputPayload) Publicize() *UpdateInputPayload {
	var pub UpdateInputPayload
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.TeamID != nil {
		pub.TeamID = ut.TeamID
	}
	if ut.UserID != nil {
		pub.UserID = ut.UserID
	}
	return &pub
}

// UpdateInputPayload user type.
type UpdateInputPayload struct {
	Name   string `form:"name" json:"name" xml:"name"`
	TeamID *int   `form:"team_id,omitempty" json:"team_id,omitempty" xml:"team_id,omitempty"`
	UserID *int   `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
}

// Validate validates the UpdateInputPayload type instance.
func (ut *UpdateInputPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// updateMemberPayload user type.
type updateMemberPayload struct {
	Role *string `form:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
}

// Validate validates the updateMemberPayload type instance.
func (ut *updateMemberPayload) Validate() (err error) {
	if ut.Role == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "role"))
	}
	return
}

// Publicize creates UpdateMemberPayload from updateMemberPayload
func (ut *updateMemberPayload) Publicize() *UpdateMemberPayload {
	var pub UpdateMemberPayload
	if ut.Role != nil {
		pub.Role = *ut.Role
	}
	return &pub
}

// UpdateMemberPayload user type.
type UpdateMemberPayload struct {
	Role string `form:"role" json:"role" xml:"role"`
}

// Validate validates the UpdateMemberPayload type instance.
func (ut *UpdateMemberPayload) Validate() (err error) {
	if ut.Role == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "role"))
	}
	return
}

// updateTwitterAccountInputPayload user type.
type updateTwitterAccountInputPayload struct {
	Name   *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	TeamID *int    `form:"team_id,omitempty" json:"team_id,omitempty" xml:"team_id,omitempty"`
	UserID *int    `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
}

// Publicize creates UpdateTwitterAccountInputPayload from updateTwitterAccountInputPayload
func (ut *updateTwitterAccountInputPayload) Publicize() *UpdateTwitterAccountInputPayload {
	var pub UpdateTwitterAccountInputPayload
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	if ut.TeamID != nil {
		pub.TeamID = ut.TeamID
	}
	if ut.UserID != nil {
		pub.UserID = ut.UserID
	}
	return &pub
}

// UpdateTwitterAccountInputPayload user type.
type UpdateTwitterAccountInputPayload struct {
	Name   *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	TeamID *int    `form:"team_id,omitempty" json:"team_id,omitempty" xml:"team_id,omitempty"`
	UserID *int    `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
}

// updateUserPayload user type.
type updateUserPayload struct {
	Bio   *string `form:"bio,omitempty" json:"bio,omitempty" xml:"bio,omitempty"`
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	Name  *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Username
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
}

// Validate validates the updateUserPayload type instance.
func (ut *updateUserPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.Username == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "username"))
	}
	if ut.Bio == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "bio"))
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Name != nil {
		if ok := goa.ValidatePattern(`\S`, *ut.Name); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.name`, *ut.Name, `\S`))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) > 256 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 256, false))
		}
	}
	if ut.Username != nil {
		if ok := goa.ValidatePattern(`^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`, *ut.Username); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.username`, *ut.Username, `^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`))
		}
	}
	if ut.Username != nil {
		if utf8.RuneCountInString(*ut.Username) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.username`, *ut.Username, utf8.RuneCountInString(*ut.Username), 40, false))
		}
	}
	return
}

// Publicize creates UpdateUserPayload from updateUserPayload
func (ut *updateUserPayload) Publicize() *UpdateUserPayload {
	var pub UpdateUserPayload
	if ut.Bio != nil {
		pub.Bio = *ut.Bio
	}
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Username != nil {
		pub.Username = *ut.Username
	}
	return &pub
}

// UpdateUserPayload user type.
type UpdateUserPayload struct {
	Bio   string `form:"bio" json:"bio" xml:"bio"`
	Email string `form:"email" json:"email" xml:"email"`
	Name  string `form:"name" json:"name" xml:"name"`
	// Username
	Username string `form:"username" json:"username" xml:"username"`
}

// Validate validates the UpdateUserPayload type instance.
func (ut *UpdateUserPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "username"))
	}
	if ut.Bio == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "bio"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, ut.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, ut.Email, goa.FormatEmail, err2))
	}
	if ok := goa.ValidatePattern(`\S`, ut.Name); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.name`, ut.Name, `\S`))
	}
	if utf8.RuneCountInString(ut.Name) > 256 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, ut.Name, utf8.RuneCountInString(ut.Name), 256, false))
	}
	if ok := goa.ValidatePattern(`^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`, ut.Username); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.username`, ut.Username, `^[\dA-Za-z]+(?:-[\dA-Za-z]+)*$`))
	}
	if utf8.RuneCountInString(ut.Username) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.username`, ut.Username, utf8.RuneCountInString(ut.Username), 40, false))
	}
	return
}