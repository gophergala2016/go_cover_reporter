
[![Travis CI](https://travis-ci.org/gophergala2016/go_cover_reporter.svg?branch=master)](https://travis-ci.org/gophergala2016/go_cover_reporter)  [![GoCoverReporter](https://ancient-beach-91563.herokuapp.com/coverage)](https://ancient-beach-91563.herokuapp.com/)

# Project GoCoverReporter

GoCoverReporter should provide access to the test coverage report data for Golang projects hosted publicly on GitHub and integrated with Travis CI. GoCoverReporter should also provide animated test coverage badge with stated test coverage percentage to be included in README.md file, and also a link to coverage statistics for the latest Travis CI run.

This software consist from two interacting parts. One is client that is running during Travis CI cycle and forwarding data from the build to the server. Server also provides for code coverage button that can be integrated into some page, such as this page you are just reading. Server also provides for landing page accessible by clicking on the badge/button.

# Demo

This page can also be used as a demo for working of GoCoverReporter. On this page, there should be visible animated button/badge with visual and quantitative representation of Golang's `go test -cover` command. If you click on that button, you should land into web page also showing same result but in visually different manner. (If there isn't a button visible on this page, that can be due to the fact that this demo is hosted free of charge :-)


## Copyright and license

Code and documentation copyright 2016 Ivan Tanaskovic. Code released under [the MIT license](https://github.com/gophergala2016/go_cover_reporter/blob/master/LICENSE)
