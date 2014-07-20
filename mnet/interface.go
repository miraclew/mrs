package mnet

// Interfaces this module depends on

type ClientCommandHandler interface {
	RecievePayload(payload *Payload)
}
