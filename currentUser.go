package sstat

import (
	"os"
	"os/user"
)

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

// GetCurrentUser returns information about the current user.
func GetCurrentUser() (*UserInfo, error) {
	var (
		userHost    *UserInfo
		currentUser *user.User
		group       *user.Group
		err         error
	)

	userHost = new(UserInfo)

	currentUser, err = user.Current()
	if err != nil {
		return nil, err
	}

	userHost.Uid = currentUser.Uid
	userHost.Gid = currentUser.Gid
	userHost.Username = currentUser.Username

	group, err = user.LookupGroupId(userHost.Gid)
	if err != nil {
		return nil, err
	}

	userHost.Group = group.Name

	userHost.Hostname, err = os.Hostname()
	if err != nil {
		return nil, err
	}

	return userHost, nil
}
