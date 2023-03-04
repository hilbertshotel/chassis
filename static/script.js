// ON PAGE LOAD
const main = async () => {
    
    const addr = "ws://127.0.0.1:8888/ws"

    // connect to ws server
    const socket = new WebSocket(addr)

    // handle ws functionality
    socket.onopen = (_) => {
        console.log("Web Socket Connection Established")
    }
    socket.onclose = (_) => {
        console.log("Web Socket Connection Terminated")
    }
    socket.onmessage = (event) => {
        // ...
    }

}

main()