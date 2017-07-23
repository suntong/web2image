////////////////////////////////////////////////////////////////////////////
// Program: web2image
// Purpose: Web to image
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import "github.com/mkideal/cli"

//==========================================================================
// Main dispatcher

func web2image(ctx *cli.Context) error {
	// ctx.JSON(ctx.RootArgv())
	// ctx.JSON(ctx.Argv())
	// fmt.Println()
	rootArgv = ctx.RootArgv().(*rootT)
	cdpScreenshot(rootArgv.Headless, ctx.Args()[0], rootArgv.CSS, ctx.Args()[1])

	return nil
}
