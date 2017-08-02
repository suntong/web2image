////////////////////////////////////////////////////////////////////////////
// Program: web2image
// Purpose: Web to image
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/mkideal/cli"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

//==========================================================================
// web2image

type rootT struct {
	cli.Helper
	Headless bool        `cli:"d,headless" usage:"use chrome-headless docker as client instead of chrome"`
	CSS      string      `cli:"*c,css" usage:"the CSS selector for the region to take the screenshot of"`
	Sleep    string      `cli:"s,sleep" usage:"time to sleep before waiting for above CSS (default: 3 seconds)" dft:"3s"`
	WaitFor  string      `cli:"w,wait-for" usage:"the CSS selector to wait for before taking the screenshot"`
	Verbose  cli.Counter `cli:"v,verbose" usage:"verbose mode (multiple -v options increase the verbosity.)"`
}

var root = &cli.Command{
	Name: "web2image",
	Desc: "Web to image\nbuilt on " + buildTime,
	Text: "Tool to take screenshot from a web page" +
		"\n\nUsage:\n  web2image OPTIONS... URL IMAGE-FILE",
	Global: true,
	Argv:   func() interface{} { return new(rootT) },
	Fn:     web2image,

	NumOption: cli.AtLeast(1),

	NumArg: cli.AtLeast(2),

	CanSubRoute: true,
}

// Template for main starts here
////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "web2image"
//          VERSION   = "0.1.0"
//          buildTime = "2017-08-02"
//  )

//  var rootArgv *rootT

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	cli.SetUsageStyle(cli.ManualStyle) // up-down, for left-right, use NormalStyle
//  	//NOTE: You can set any writer implements io.Writer
//  	// default writer is os.Stdout
//  	if err := cli.Root(root,).Run(os.Args[1:]); err != nil {
//  		fmt.Fprintln(os.Stderr, err)
//  	}
//  	fmt.Println("")
//  }

// Template for main dispatcher starts here
//==========================================================================
// Main dispatcher

//  func web2image(ctx *cli.Context) error {
//  	ctx.JSON(ctx.RootArgv())
//  	ctx.JSON(ctx.Argv())
//  	fmt.Println()

//  	return nil
//  }

// Template for CLI handling starts here
