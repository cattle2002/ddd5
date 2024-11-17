package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8084", "http service address")

var upgrader = websocket.Upgrader{} // use default options

const (
	login              = "login"
	loginRet           = "loginRet"
	monitor            = "monitor"
	monitorRet         = "monitorRet"
	alarm              = "alarm"
	alarmRet           = "alarmRet"
	ticket             = "ticket"
	ticketRet          = "ticketRet"
	packageUpload      = "packageUpload"
	packageUploadRet   = "packageUploadRet"
	packageAllocate    = "packageAllocate"
	packageAllocateRet = "packageAllocateRet"
	release            = "release"
	releaseRet         = "releaseRet"
	operationAudit     = "operationAudit"
	operationAuditRet  = "operationAuditRet"
	control            = "control"
	controlRet         = "controlRet"
	autoUpgrade        = "autoUpgrade"
	autoUpgradeRet     = "autoUpgradeRet"
	adminModify        = "adminModify"
	adminModifyRet     = "adminModifyRet"
	push               = "push"
	pushRet            = "pushRet"
	orderPush          = "orderPush"
	orderPushRet       = "orderPushRet"
)

type CommonReq struct {
	Command string `json:"Command"`
	Id      int64  `json:"Id"`
}

func handle(message []byte, conn *websocket.Conn) {
	var req CommonReq
	err := json.Unmarshal(message, &req)
	if err != nil {
		log.Println("WebSocket Unmarshal error:", err)
		return
	}
	if req.Command == login {

	}

}
func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;

    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
        output.scroll(0, output.scrollHeight);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
</td></tr></table>
</body>
</html>
`))
