package sstat

import (
	"os"
	"os/user"
)

// UserInfo contains user information.
type UserInfo struct {
	// Uid is the user ID.
	// On POSIX systems, this is a decimal number representing the uid.
	// On Windows, this is a security identifier (SID) in a string format.
	// On Plan 9, this is the contents of /dev/user.
	Uid string

	// Gid is the primary group ID.
	// On POSIX systems, this is a decimal number representing the gid.
	// On Windows, this is a SID in a string format.
	// On Plan 9, this is the contents of /dev/user.
	Gid string

	// Username is the login name.
	Username string

	// Group is the primary group name.
	Group string

	// Hostname is the host name reported by the kernel.
	Hostname string
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
		Uid:      account.Uid,
		Gid:      account.Gid,
		Username: account.Username,
		Group:    group.Name,
	}

	userInfo.Hostname, err = os.Hostname()
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
