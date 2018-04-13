# Grellyd_Site
## Golang Website

This my website written in Golang. It was originally based off of the [wiki tutorial](https://golang.org/doc/articles/wiki/) for an understanding of routers etc, but has since expanded.

It is deployed on AWS via ElasticBeanstalk in a Docker Container, per the local dockerfile.

## Future Steps:
1. Append a blog section (either homemade or via [Hugo](https://gohugo.io/)
    - If hugo, it will be a custom theme spawned from a mixture of [paperback](https://github.com/damiencaselli/paperback), [hugo-paper](https://github.com/nanxiaobei/hugo-paper), and [hugo-xmin](https://github.com/yihui/hugo-xmin).
    - All the content for the site will reside in the `content` folder, with a distinction between static pages and blog posts. 
    - There will be a few `layout`s which specifiy the page type, similar to my templates.
    - Styling will probably be the largest obstacle.
2. S3 Bucket for resumes on site with links
