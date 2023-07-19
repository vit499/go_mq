import { makeAutoObservable, runInAction } from "mobx";
import mq from "./Mq";

class WsStore {
  //_connectStatus;
  constructor() {
    this._listMessage = [];
    this.Message = "-";
    this._connectStatus = "Closed";
    this.username = "a";
    this.savedMes = "";
    this.sock = null;
    this.url = "";
    this.eff = false;
    this.url = ""; //process.env.REACT_APP_API_URL || "ws://localhost:8080/ws";
    makeAutoObservable(this, {});
  }

  wsConnect(url) {
    //console.log(`_connectStatus=${this._connectStatus}`);
    if (this._connectStatus !== "Closed") {
      //return;
    }
    this._connectStatus = "Connecting";

    this.sock = new WebSocket(this.url);

    this.sock.onopen = () => {
      runInAction(() => {
        this._connectStatus = "Connected";
        console.log("connected");
        this.eff = false;
        //this.sendInfo();
        if (this.savedMes !== "") {
          this.sendMessage(this.savedMes);
          this.savedMes = "";
        }
      });
    };

    /*
    const wsRecMes = {
      username: this.username,
      topic: "ab@c.ru/5555/devsend/cp",
      message: "0001225577",
      group: "mqtt"
    }
    */
    this.sock.onmessage = (event) => {
      // runInAction(() => {
      //   const message = JSON.parse(event.data);
      //   //this.Message = message.message;
      //   //console.log(JSON.stringify(message, null, 2));
      // });
      const src = JSON.parse(event.data);
      mq.ReceiveMesFromWs(src);
    };
    this.sock.onclose = () => {
      runInAction(() => {
        this._connectStatus = "Closed";
        this.Message = "-";
        //console.log("ws closed");
      });
    };
    this.sock.onerror = () => {
      runInAction(() => {
        this._connectStatus = "Closed";
        this.Message = "-";
        //console.log("ws error");
      });
    };
  }

  sendMessage = async (val) => {
    if (this._connectStatus !== "Connected") {
      this.savedMes = val;
      this.wsConnect();
      return;
    }
    const message = {
      username: this.username,
      message: val,
      id: Date.now(),
      event: "message",
    };
    this.sock.send(JSON.stringify(message));
  };
  sendInfo = async () => {
    const message = {
      username: this.username,
      id: Date.now(),
      event: "connection",
    };
    this.sock.send(JSON.stringify(message));
  };

  Disconnect() {
    //console.log(`dis eff=${this.eff}`);
    if (this.eff) return;
    if (this._connectStatus === "Closed") {
      return;
    }
    this.sock.close();
  }

  Init() {
    if (this.eff) return;
    this.eff = true;
    if (this._connectStatus !== "Closed") {
      return;
    }
    this.username = Date.now();
    //this.url = process.env.REACT_APP_API_URL || "ws://localhost:3015";
    //const url = `ws://localhost:3015`;
    //this.url = process.env.REACT_APP_API_URL || "ws://localhost:8080/ws";
    this.url = "ws://localhost:3100/ws";
    console.log(`url=${this.url}`);
    this.wsConnect();
  }
}

const wsStore = new WsStore();

export default wsStore;
