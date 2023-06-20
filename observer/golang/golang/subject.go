package main

type Subject interface {
	register(observer Observer)
	deregister(obs Observer)
	notifyAll()
}
