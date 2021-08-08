package polysocket

type State uint8

const (
	// STInvalid Zero value
	STInvalid State = 0
	// STError Client encountered kind of unrecoverable error.
	STError State = 1
	// STClosed Client was closed by calling Close()
	STClosed State = 2
	// STDisconnected Client::Disconnect() was called.
	STDisconnected State = 3
	// STUnconnected Client object was just initialized and no operations have
	// been performed.
	STUnconnected State = 4
	// STConnecting Client between when Connect() was called and
	// the server sending a confirmation "connected" message.
	STConnecting State = 5
	// STConnected Client is connected but not yet authenticated
	STConnected State = 6
	// STReady Client is connected and ready to subscribe.
	STReady State = 7
	// STReconnecting Client is in the process of reconnecting.
	STReconnecting State = 8
)

func (cs State) String() string {
	switch cs {
	case STError:
		return "STError"
	case STUnconnected:
		return "STUnconnected"
	case STClosed:
		return "STClosed"
	case STConnecting:
		return "STConnecting"
	case STConnected:
		return "STConnected"
	case STReady:
		return "STReady"
	case STReconnecting:
		return "STReconnecting"
	default:
		return "STInvalid"
	}
}
