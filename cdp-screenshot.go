package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	cdp "github.com/knq/chromedp"
	"github.com/knq/chromedp/client"
)

func cdpScreenshot(headless bool, url, css, imgf string) {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	var c *cdp.CDP
	if headless {
		// create chrome-headless instance
		c, err = cdp.New(ctxt, cdp.WithTargets(client.New().WatchPageTargets(ctxt)))
	} else {
		// create chrome instance
		c, err = cdp.New(ctxt)
	}
	abortOn("Create chrome instance", err)
	if rootOpts.Verbose >= 1 {
		err := cdp.WithErrorf(log.Printf)(c)
		abortOn("Turn on logging", err)
	}

	// run task list
	var buf []byte
	err = c.Run(ctxt, screenshot(url, css, &buf))
	if err != nil {
		log.Fatal(err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(imgf, buf, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func screenshot(urlstr, sel string, res *[]byte) cdp.Tasks {
	s, err := time.ParseDuration(rootOpts.Sleep)
	abortOn("Parsing sleep time", err)

	t := cdp.Tasks{
		cdp.Navigate(urlstr),
		cdp.Sleep(s),
		cdp.WaitVisible(sel, cdp.ByQuery),
		cdp.Screenshot(sel, res, cdp.ByQuery),
	}
	// if len(Opts.WaitFor) != 0 {
	//   cdp.WaitVisible(Opts.WaitFor, cdp.ByQuery),
	// }
	return t
}
