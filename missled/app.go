package main

import (
	"github.com/miraclew/mrs/missle"
	"github.com/miraclew/mrs/mnet"
	"github.com/miraclew/mrs/push"
	"log"
	"net"
	"sync"
)

type App struct {
	options      *AppOptions
	tcpAddr      *net.TCPAddr
	httpAddr     *net.TCPAddr
	tcpListener  net.Listener
	httpListener net.Listener
	waitGroup    sync.WaitGroup
	exitChan     chan int
}

type AppOptions struct {
}

func NewApp(options *AppOptions) *App {
	a := &App{
		options:  options,
		exitChan: make(chan int),
	}

	return a
}

func NewAppOptions() *AppOptions {
	options := &AppOptions{}

	return options
}

func (a *App) Main() {
	missle.SetDSN(DSN)
	pusher := &push.Pusher{}
	game := missle.GetGame()

	game.HandlePush(pusher)
	pusher.HandleConnection(game)

	server := mnet.NewServer()
	manager := mnet.NewManager(server, nil)

	tcpListener, err := net.Listen("tcp", a.tcpAddr.String())
	if err != nil {
		log.Fatalf("FATAL: listen (%s) failed - %s", a.tcpAddr, err.Error())
	}
	a.tcpListener = tcpListener

	httpListener, err := net.Listen("tcp", a.httpAddr.String())
	if err != nil {
		log.Fatalf("FATAL: listen (%s) failed - %s", a.tcpAddr, err.Error())
	}
	a.httpListener = httpListener

	a.waitGroup.Add(1)
	go func() {
		httpServe(httpListener)
		a.waitGroup.Done()
	}()

	a.waitGroup.Add(1)
	go func() {
		// pusher.Serve(tcpListener)
		manager.Serve(tcpListener)
		a.waitGroup.Done()
	}()
}

func (a *App) Exit() {
	if a.tcpListener != nil {
		a.tcpListener.Close()
	}

	if a.httpListener != nil {
		a.httpListener.Close()
	}

	close(a.exitChan)
	a.waitGroup.Wait()
}
