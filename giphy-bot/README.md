# Giphy Bot

A giphy bot which is built with go-lark/gin/gorm as an real world example.

This bot is not built for production use.

![Giphy hello world](/giphy-bot/assets/giphy-hello-world.gif)

## Usage

```plain
/i: id
/s: search ...keywords
/r: random ...keywords
/t: trending [sequence]
```

* Giphy: Trending
  * `/t`: Trending
  * `/t 5`: Trending, the 5th one
* Giphy by ID
  * `/i 2A5wXdxrWghhSvvFFb`: Send Giphy by ID
* Giphy by Search
  * `/s hello world`: Search Giphy with keywords
  * HINT: `/s` can be ignored. Just type your keywords.
* Giphy by Random
  * `/r hello world`: Search Giphy with keywords, and return a random one

## FAQ

* Why Giphy bot is so slow?
  * Because it's really slow to fetch images from giphy server from China.
* Why does Giphy bot always reply me twice?
  * Lark Platform will retry if the bot does not respond within 1 second. Unluckily, Giphy bot is slow as aforementioned reason.
  * Solved with asynchronous way on Nov 28, 2019.
* Why Giphy is not available?
  * ![Giphy is not available](/giphy-bot/assets/giphy-not-available.gif)
  * Technical problems with giphy.com or image size is larger than 10MiB.
