package fipsvc

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Hook func()

var (
	cleanupHooks []Hook
	reloadHooks  []Hook
)

func AddCleanupHooks(hooks ...Hook) {
	cleanupHooks = append(cleanupHooks, hooks...)
}

func AddReloadHooks(hooks ...Hook) {
	reloadHooks = append(reloadHooks, hooks...)
}

func Start() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	go func() {
		<-c
		execHooks(cleanupHooks)
		os.Exit(1)

	}()

	r := make(chan os.Signal, 1)
	signal.Notify(r, syscall.SIGHUP)
	go func() {
		<-r
		execHooks(reloadHooks)
		log.Println("Service reloaded")
	}()
}

func execHooks(hooks []Hook) {
	for _, f := range hooks {
		f()
	}
}
