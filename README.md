# godu

Find the files that are taking up your space.

Tired of looking like a noob with [Disk Inventory X](http://www.derlien.com/) or SpaceMonger? Do you want something that
* can do the job
* works in terminal
* makes you look cool
* is written in Golang
* you can contribute to

??

Well then **look no more** and try out the godo.

## Installation
```
go get -u github.com/tomasvik/godu
```

## Usage
```
godo ~
godo folder-that-is-too-big
```

Once the folder is crawled (can take up to few minutes), you move around by selecting numbers (moving deeper in the structure) or you just press enter (moving up in the structure)

Exit with CTRL+C