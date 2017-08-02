
# {{.Name}}

{{render "license/shields" . "License" "MIT"}}
{{template "badge/godoc" .}}
{{template "badge/goreport" .}}
{{template "badge/travis" .}}

## {{toc 5}}

## {{.Name}} - Web to image

The `web2image` will take a screenshot of a given web page.

# Usage

### $ {{exec "web2image" | color "sh"}}

### chrome-headless

The following examples assumes that you've [setup chrome-headless](https://github.com/knq/chromedp/blob/master/examples/headless/README.md) already. Else, take a look at [Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/) on how to start a local chrome with the remote-debugging-port opening at `9222`.

# Examples

## Film info from cmore.se

    web2image -d -c '#main-wrapper' 'http://www.cmore.se/film/3643033-deadpool' example-deadpool.png

and the result is:

![deadpool](example-deadpool.png "deadpool.png")

## US market indexes chart

    web2image -d -s '5s' -c 'div > section.market-indexes-component > div.market-indexes-container' http://www.morningstar.com/ example-US-market-indexes.png

and the result is:

![US-market-indexes](example-US-market-indexes.png "US-market-indexes.png")

# Installation

```
go get github.com/suntong/web2image
```

Debian package will be provided shortly.

## Author

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

All patches welcome. 
