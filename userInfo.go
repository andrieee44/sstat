package sstat

import (
	"os"
	"os/user"
)

// UserInfo contains user information.
type UserInfo struct {
	uid      string
	gid      string
	username string
	group    string
	hostname string
}

// Uid reports the user ID.
// On POSIX systems, this is a decimal number representing the uid.
// On Windows, this is a security identifier (SID) in a string format.
// On Plan 9, this is the contents of /dev/user.
func (info *UserInfo) Uid() string {
	return info.uid
}

// Gid reports the primary group ID.
// On POSIX systems, this is a decimal number representing the gid.
// On Windows, this is a SID in a string format.
// On Plan 9, this is the contents of /dev/user.
func (info *UserInfo) Gid() string {
	return info.gid
}

// Username reports the login name.
func (info *UserInfo) Username() string {
	return info.username
}

// Group reports the primary group name.
func (info *UserInfo) Group() string {
	return info.group
}

// Hostname reports the host name reported by the kernel.
func (info *UserInfo) Hostname() string {
	return info.hostname
}

func mkUserInfo(account *user.User) (*UserInfo, error) {
	var (
		userInfo *UserInfo
		group    *user.Group
		err      error
	)

	group, err = user.LookupGroupId(account.Gid)
	if err != nil {
		return nil, err
	}

	userInfo = &UserInfo{
		uid:      account.Uid,
		gid:      account.Gid,
		username: account.Username,
		group:    group.Name,
	}

	userInfo.hostname, err = os.Hostname()
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

// LookupUser looks up a user by username. If the user cannot be found,
// the returned error is of type UnknownUserIdError.
func LookupUser(username string) (*UserInfo, error) {
	var (
		account *user.User
		err     error
	)

	account, err = user.Lookup(username)
	if err != nil {
		return nil, err
	}

	return mkUserInfo(account)
}

// LookupUserId looks up a user by userid. If the user cannot be found,
// the returned error is of type UnknownUserIdError.
func LookupUserId(uid string) (*UserInfo, error) {
	var (
		account *user.User
		err     error
	)

	account, err = user.LookupId(uid)
	if err != nil {
		return nil, err
	}

	return mkUserInfo(account)
}

// CurrentUser returns information about the current user.
func CurrentUser() (*UserInfo, error) {
	var (
		account *user.User
		err     error
	)

	account, err = user.Current()
	if err != nil {
		return nil, err
	}

	return mkUserInfo(account)
}
