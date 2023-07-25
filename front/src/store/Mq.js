import { makeAutoObservable, runInAction } from "mobx";
//import mqtt from "mqtt";
import * as mqtt from "mqtt/dist/mqtt.min";
import descrStore from "./DescrStore";
import hostStore from "./HostStore";
import temperStore from "./TemperStore";

class Mq {
  _connectStatus;
  constructor() {
    hostStore.getHostFromStorage();
    descrStore.getDescrFromStorage();
    this._client = null;
    this._isSub = false;
    this._listMessage = [];
    this._connectStatus = "Connect";
    this.updTopics();
    makeAutoObservable(this, {});
  }

  // get numObject() {
  //   return this._numObject;
  // }
  // get numObject2() {
  //   return this._numObject2;
  // }
  get topicPub() {
    return this._topicPub;
  }
  // get connectStatus() {
  //   return this._connectStatus;
  // }

  setConnectStatus(s) {
    this._connectStatus = s;
  }
  updTopics() {
    this._topicSub = hostStore.login + "/" + hostStore.obj1 + "/devsend/#";
    this._topicPub = hostStore.login + "/" + hostStore.obj1 + "/devrec/control";
    this._topicSub2 = hostStore.login + "/" + hostStore.obj2 + "/devsend/#";
    this._topicSub3 = hostStore.login + "/" + hostStore.obj3 + "/devsend/#";
    this._topicPub2 =
      hostStore.login + "/" + hostStore.obj2 + "/devrec/control";
    this._topicPub3 =
      hostStore.login + "/" + hostStore.obj3 + "/devrec/control";
  }
  formTopicPub(indObj) {
    const topic =
      hostStore.login +
      "/" +
      temperStore._nvobj[indObj].nobj +
      "/devrec/control";
    return topic;
  }

  mqttConnect(host, mqttOption) {
    if (this._connectStatus !== "Connect") {
      return;
    }
    console.log("mq try connect");
    this._connectStatus = "Connecting";
    this._client = mqtt.connect(host, mqttOption);

    if (this._client) {
      this._client.on("connect", () => {
        runInAction(() => {
          this._connectStatus = "Connected";
          console.log("connected");
          // authStore.setAuth(true);
          this._isSub = false;
          this.mqttSub();
        });
      });
      this._client.on("error", (err) => {
        runInAction(() => {
          console.error("Connection error: ", err);
          this._client.end();
          this._isSub = false;
          this._client = null;
          this._connectStatus = "Connect";
        });
      });
      this._client.on("reconnect", () => {
        runInAction(() => {
          this._connectStatus = "Reconnecting";
          this._isSub = false;
        });
      });
      this._client.on("message", (topic, message) => {
        // runInAction(() => {
        //   // const id = Math.random().toString().substr(2, 10);
        //   // const p = { id, topic, message: message.toString() };
        //   // //this._listMessage.push(p);
        //   // temperStore.recMes(p);
        //   // // console.log("mes:", JSON.stringify(this._listMessage, null, 2));
        // });
        this.ReceiveMesBytes(topic, message);
      });
    }
  }

  ReceiveMesBytes(topic, mesBytes) {
    message = mesBytes.toString();
    this.ReceiveMes(topic, message);
  }
  ReceiveMes(topic, message) {
    runInAction(() => {
      temperStore.recMes(topic, message);
    });
  }

  ReceiveMesFromWs(src) {
    let topic = "";
    let message = "";
    let username = "";
    let group = "";
    try {
      username = src.username;
      topic = src.topic;
      message = src.message;
      group = src.group;
    } catch (e) {
      console.log(e);
    }
    //console.log(`username: ${username}, topic: ${topic}, group: ${group}`);
    if (username != hostStore.login) return;
    if (group == "mqtt") {
      //this.ReceiveMes(topic, message);
    } else if (group == "json") {
      runInAction(() => {
        temperStore.recMesJson(topic, message);
      });
    }
  }
  mqttDisconnect() {
    if (this._client) {
      console.log("disconnect");
      this._client.end(() => {
        runInAction(() => {
          this._connectStatus = "Connect";
          this._isSub = false;
          this._client = null;
          temperStore.clearAll();
        });
      });
    }
  }
  mqttPublish(context) {
    if (this._client) {
      const qos = 1;
      const { indObj, payload } = context;
      temperStore.clear(indObj);
      const topicPub = this.formTopicPub(indObj);
      this._client.publish(topicPub, payload, { qos }, (error) => {
        runInAction(() => {
          if (error) {
            console.log("Publish error: ", error);
          } else {
            console.log(`${topicPub} ${payload} `);
          }
        });
      });
    }
  }
  mqttPublishAll(context) {
    if (this._client) {
      // ab@m.ru/0802/devrec/control
      // this._topicPub =
      //  this._login + "/" + this._numObject + "/devrec/control";
      const qos = 1;
      const { payload } = context;
      temperStore.clearAll();
      const topicPub1 = this.formTopicPub(0);
      this._client.publish(topicPub1, payload, { qos }, (error) => {
        runInAction(() => {
          if (error) {
            console.log("Publish error: ", error);
          } else {
            console.log(`${topicPub1} ${payload} `);
          }
        });
      });
      const topicPub2 = this.formTopicPub(1);
      this._client.publish(topicPub2, payload, { qos }, (error) => {
        runInAction(() => {
          if (error) {
            console.log("Publish2 error: ", error);
          } else {
            console.log(`${topicPub2} ${payload} `);
          }
        });
      });
      const topicPub3 = this.formTopicPub(2);
      this._client.publish(topicPub3, payload, { qos }, (error) => {
        runInAction(() => {
          if (error) {
            console.log("Publish3 error: ", error);
          } else {
            console.log(`${topicPub3} ${payload} `);
          }
        });
      });
    }
  }

  mqttSub() {
    if (this._client && !this._isSub) {
      // ab@m.ru/0802/devsend/#
      // this._topicSub = this._login + "/" + this._numObject + "/devsend/#";
      const qos = 1;
      //const { topic, qos } = subscription;
      this._client.subscribe(
        [this._topicSub, this._topicSub2, this._topicSub3],
        { qos },
        (error) => {
          if (error) {
            runInAction(() => {
              console.log("Subscribe to topics error", error);
            });
          } else {
            runInAction(() => {
              console.log("sub ok");
              this._isSub = true;
              this.mqttPublishAll({ payload: "reqfout" });
            });
          }
        }
      );
    }
  }

  mqttUnSub() {
    if (this._client) {
      //const { topic } = subscription;
      this._client.unsubscribe(
        [this._topicSub, this._topicSub2, this._topicSub3],
        (error) => {
          runInAction(() => {
            if (error) {
              console.log("Unsubscribe error", error);
              return;
            }
            this._isSub = false;
          });
        }
      );
    }
  }

  mqTryConnect() {
    if (this._connectStatus !== "Connect") {
      return;
    }
    const url = `ws://${hostStore.host}:${hostStore.port}`;
    const clientId = `112${Math.random().toString(16).substr(2, 12)}`;
    const options = {
      keepalive: 10,
      protocolId: "MQTT",
      protocolVersion: 4,
      clean: true,
      reconnectPeriod: 0,
      connectTimeout: 10 * 1000,
    };
    options.clientId = clientId;
    options.username = hostStore.login;
    options.password = hostStore.password;
    this.mqttConnect(url, options);
  }
}

const mq = new Mq();

export default mq;
