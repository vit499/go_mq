import { makeAutoObservable, runInAction } from "mobx";
import hostStore from "./HostStore";
import mq from "./Mq";
import temperStore from "./TemperStore";

class WsStore {
  //_connectStatus;
  constructor() {
    this._listMessage = [];
    this.pass = "";
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
        this.sendInfo();
        // if (this.savedMes !== "") {
        //   this.sendMessage(this.savedMes);
        //   this.savedMes = "";
        // }
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

  WsPublish = async (val) => {
    const { indObj, payload } = val;
    temperStore.clear(indObj);
    const topicPub = mq.formTopicPub(indObj);

    if (this._connectStatus !== "Connected") {
      //this.savedMes = val;
      this.wsConnect();
      return;
    }
    const mes = {
      username: this.username,
      topic: topicPub,
      message: payload,
      group: "command",
      pass: this.pass,
    };
    this.sock.send(JSON.stringify(mes));
  };
  sendInfo = async () => {
    // Username: s[0],
    // Topic:    s[1],
    // Message:  s[2],
    // Group:    "mqtt",
    const mes = {
      username: this.username,
      topic: "-",
      message: "-",
      group: "connection",
      pass: this.pass,
    };
    this.sock.send(JSON.stringify(mes));
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
    this.username = hostStore.login; //Date.now();
    this.pass = hostStore.password;
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
