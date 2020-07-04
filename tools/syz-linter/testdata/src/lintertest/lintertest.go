// Copyright 2020 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package lintertest

/* some comment */ // want "Use C-style comments // instead of /* */"
var comment = 1    /* some comment */ // want "Use C-style comments // instead of /* */"

func stringComparison() {
	str := ""
	if len(str) == 0 { // want "Compare string with \"\", don't compare len with 0"
	}
	if 0 != len(str) { // want "Compare string with \"\", don't compare len with 0"
	}
	if len(returnString()+"foo") > 0 { // want "Compare string with \"\", don't compare len with 0"
	}
}

func returnString() string { return "foo" }

//
// One space.
//  Two spaces.
//	One tab.
//		Two tabs.
//No space.					// want "Use either //<one-or-more-spaces>comment or //<one-or-more-tabs>comment format for comments"
//	  Tab and spaces.			// want "Use either //<one-or-more-spaces>comment or //<one-or-more-tabs>comment format for comments"
// 	Space and tab.				// want "Use either //<one-or-more-spaces>comment or //<one-or-more-tabs>comment format for comments"
func checkCommentSpace() {
}