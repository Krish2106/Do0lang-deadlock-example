# Golang-deadlock-example
An example of deadlock in Golang.

A deadlock coudl arise due to multiple reasons, discussed below is a common scenario where the unbuffered channel waits infinetely to recieve signal from a go-routine which leads to a deadlock.
