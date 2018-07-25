# Hi!
## Welcome to the personal website of Graham L. Brown
### Backend Developer, Backpacker, and Student

This page is under construction. Please excuse the mess. 

In the meantime, connect with me!
 * [Github](https://github.com/grellyd)
 * [LinkedIn](https://www.linkedin.com/in/grellyd/)
 * [Twitter](https://twitter.com/grellyd/)

title: "{{ .Name | replaceRE "^([0-9_]{11})([a-zA-Z_])" "$2" | replaceRE  "_" " " | title }}"
title: "{{ .Name | slicestr 11 | replaceRE  "_" " " | title }}"
