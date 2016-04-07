// Copyright 2010-2016 Joubin Houshyar.  All rights reserved.
// Use of this source code is governed by a 2-clause BSD
// license that can be found in the LICENSE file.

package goerror_test

import (
	"fmt"
	"goerror"
)

// Let's define a few canonical goerror
var (
	IllegalArgument = goerror.Define("IllegalArgument")
	IllegalState    = goerror.Define("IllegalState")
	Bug             = goerror.Define("BUG")
)

// Example defining, returning, and checking goerror
func ExampleError() {

	user := "theuser"
	oldpw := "old-secret"
	newpw := "new-secret"

	if e := ChangePassword(user, oldpw, newpw); e != nil {
		switch typ := goerror.TypeOf(e); {
		case typ.Is(IllegalArgument): /* handle it */
		case typ.Is(IllegalState): /* handle it */
		default: /* this violates the API contract - must be a bug */
			panic(Bug(fmt.Sprintf("unexpected error %v returned by ChangePassword()", e)))
		}
	}
}

// (Example function that returns categorical goerror.)
// Change the user's password.
//
// returns IllegalArgument for any nil input;
// IllegalState if user not logged in.
func ChangePassword(user, oldPassword, newPassword string) error {
	// assert args
	if user == "" {
		return IllegalArgument("user is nil")
	}
	if oldPassword == "" {
		return IllegalArgument("oldpassword is nil")
	}
	if newPassword == "" {
		return IllegalArgument("newPassword is nil")
	}

	// user must be already logged in to change passwords
	// (it's just an example ;-)
	if !UserLoggedIn(user) {
		return IllegalState("user must be logged in to change pw")
	}

	// ...

	return nil
}

func UserLoggedIn(user string) bool {
	// ...
	return false
}
