package network

import(
	"net"
)

const(
	sleepduration = 1000 //interval between alivemessages given in milliseconds
	toleratedLosses = 4

	isAlive = 1
	dead = 0
)

var(
	localIP = getMyIP()
	broadcast = "235.241.187.255" //må se nærmere på adressen
	
	UDPport = "8769" // randomly chosen ports
	TCPport = " 8770"

)

var(
	updateTCPmap chan TCPconnection // new TCP connections are shared over this channel
	newIPchan chan string // new IPs broadcasting UDP are shared here
	isDeadchan chan string // when UDP module detects that someone is dead, their IP is transmitted here
	isAlivechan chan int // for internal use in UDP module. When new ping is received, input to this channel resets death timer
	
	giveMeCurrentMap chan bool
	getCurrentMap chan map[string]net.Conn
	
	giveMeConn chan string
	getSingleConn chan net.Conn
)

type TCPconnection struct { // inputs to map containing active TCP connections are of this type. IP is key, socket is content
	socket net.Conn
	IP string
}

type commChannels struct { // collection of channels used for TCP communication
	sendToAll chan message
	sendToOne chan message
	messageReceivedchan chan message
}

type message struct { // messages sent over TCP are converted to this type, before being transmitted over channels
	IP string
	content []byte
