// Copyright (c) 2015 Pagoda Box Inc
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v.
// 2.0. If a copy of the MPL was not distributed with this file, You can obtain one
// at http://mozilla.org/MPL/2.0/.
//

package stylish

import (
  "fmt"
  // "reflect"
  "strings"

  "github.com/mitchellh/go-wordwrap"
)

// Header styles and prints a header as outlined at:
// http://nanodocs.gopagoda.io/engines/style-guide#header
//
// Usage:
// Header "i am a header"
//
// Output:
// :::::::::::::::::::::::::: I AM A HEADER :::::::::::::::::::::::::
func Header(header string) {

  maxLen := 70
  subLen := len(fmt.Sprintf(" %v ", header))

  leftLen := (maxLen - subLen)/2 + (maxLen - subLen)%2
  rightLen := (maxLen - subLen)/2

  // print the header, inserting a ':' (colon) 'n' times, where 'n' is the number
  // remaining after subtracting subLen (number of 'reserved' characters) from
  // maxLen (maximum number of allowed characters)
  output := fmt.Sprintf("%v %v %v", strings.Repeat(":", leftLen), strings.ToUpper(header), strings.Repeat(":", rightLen))

  fmt.Printf(`
%v
`, output)
}

// ProcessStart styles and prints a 'child process' as outlined at:
// http://nanodocs.gopagoda.io/engines/style-guide#child-process
//
// Usage:
// ProcessStart "i am a process"
//
// Output:
// I AM A PROCESS :::::::::::::::::::::::::::::::::::::::::::::::: =>
func ProcessStart(process string) {

  maxLen := 70
  subLen := len(fmt.Sprintf("%v  =>", process))

  // print the process, inserting a ':' (colon) 'n' times, where 'n' is the number
  // remaining after subtracting subLen (number of 'reserved' characters) from
  // maxLen (maximum number of allowed characters)
  output := fmt.Sprintf("%v %v =>", strings.ToUpper(process), strings.Repeat(":", (maxLen-subLen)))

  fmt.Printf(`
%v
`, output)
}

// ProcessEnd styles and prints a 'child process' as outlined at:
// http://nanodocs.gopagoda.io/engines/style-guide#child-process
//
// Usage:
// ProcessEnd "i am a process"
//
// Output:
// <= :::::::::::::::::::::::::::::::::::::::::::: END I AM A PROCESS
func ProcessEnd(process string) {

  maxLen := 70
  subLen := len(fmt.Sprintf("<=  END %v", process))

  // print the process, inserting a ':' (colon) 'n' times, where 'n' is the number
  // remaining after subtracting subLen (number of 'reserved' characters) from
  // maxLen (maximum number of allowed characters)
  output := fmt.Sprintf("<= %v END %v", strings.Repeat(":", (maxLen-subLen)), strings.ToUpper(process))

  fmt.Printf(`
%v
`, output)
}

// SubTask styles and prints a 'sub task' as outlined at:
// http://nanodocs.gopagoda.io/engines/style-guide#sub-tasks
//
// Usage:
// SubTask "i am a sub task"
//
// Output:
// ::::::::: I AM A SUB TASK
func SubTaskStart(task string) {
  fmt.Printf(`
::::::::: %v
`, strings.ToUpper(task))
}

// SubTaskSuccess styles and prints a footer to a successful subtask
//
// Usage:
// SubTaskSuccess
//
// Output:
// <<<<<<<<< [√] SUCCESS
func SubTaskSuccess() {
  fmt.Printf("\n<<<<<<<<< [√] SUCCESS\n")
}

// SubTaskFail styles and prints a footer to a failed subtask
//
// Usage:
// SubTaskFail
//
// Output:
// <<<<<<<<< [!] FAILED
func SubTaskFail() {
  fmt.Printf("\n<<<<<<<<< [!] FAILED\n")
}

// Bullet styles and prints a message as outlined at:
// http://nanodocs.gopagoda.io/engines/style-guide#bullet-points
//
// Usage:
// Bullet "i am a bullet"
// Bullet []string{"we", "are", "many", "bullets"}
//
// Output:
// +> i am a bullet
//
// +> we
// +> are
// +> many
// +> bullets
func Bullet(v interface{}) {
  switch t := v.(type) {

  // if it's a slice of strings, iterate over each one printing them individually...
  case []string:
    for i := range t {
      fmt.Printf("+> %v\n", t[i])
    }

  // otherwise just print the value
  default:
    fmt.Printf("+> %v\n", v)
  }
}

// Warning styles and prints a message as outlined at:
// http://nanodocs.gopagoda.io/engines/style-guide#warning
//
// Usage:
// Warning "You just bought Hot Pockets!"
//
// Output:
// -----------------------------  WARNING  -----------------------------
// You just bought Hot Pockets!
func Warning(body string) {
  fmt.Printf(`
-----------------------------  WARNING  -----------------------------
%v
`, wordwrap.WrapString(body, 70))
}

// Error styles and prints a message as outlined at:
// http://nanodocs.gopagoda.io/engines/style-guide#fatal_errors
//
// Usage:
// Error "nuclear launch detected", "All your base are belong to me"
//
// Output:
// ! NUCLEAR LAUNCH DETECTED !
//
// All your base are belong to me
func Error(heading, body string) {
  fmt.Printf(`
! %v !

%v
`, strings.ToUpper(heading), wordwrap.WrapString(body, 70))
}
