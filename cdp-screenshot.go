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
		c, err = cdp.New(ctxt, cdp.WithTargets(client.New().WatchPageTargets(ctxt)), cdp.WithErrorf(log.Printf))
	} else {
		// create chrome instance
		c, err = cdp.New(ctxt)
	}
	if err != nil {
		log.Fatal(err)
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
	return cdp.Tasks{
		cdp.Navigate(urlstr),
		cdp.Sleep(2 * time.Second),
		cdp.WaitVisible(sel, cdp.ByID),
		cdp.Screenshot(sel, res, cdp.ByID),
	}
}
