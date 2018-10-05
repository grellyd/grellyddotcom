# Grellyd_Site

This my website at [grellyd.com](https://grellyd.com) using a Go server to server content generated using [Hugo](https://gohugo.io) with my custom theme [Kubo](https://github.com/grellyd/kubo_hugo).

The server was originally based off of the [wiki tutorial](https://golang.org/doc/articles/wiki/) for an understanding of routers etc, but has since expanded. As the site is not served simply by `hugo serve`, there is the ability to easily add other projects and services.

It is deployed on AWS via ElasticBeanstalk in a Docker Container, which gets rebuilt every 'eb deploy' with a fresh pull from Github.

## Future Steps:
1. Append a blog section (via [Hugo](https://gohugo.io/)
2. Blog about building the Kubo Theme
    - Originally a mixture of [paperback](https://github.com/damiencaselli/paperback), [hugo-paper](https://github.com/nanxiaobei/hugo-paper), and [hugo-xmin](https://github.com/yihui/hugo-xmin).
    - Now its own beast.
    - Lessons Learned
    - What could be changed to improve template writing for first time authors.
        Within the GA template, auto remove localhost
        Clear and concise instructions re how to get static pages with Hugo
        Partial org and baseof

## Flow
- Markdown content lives in content/*
- Generated html lives in public/*
- Hugo takes the markdown files and generates the public folder.
- Go server serves the static files out of public/*
