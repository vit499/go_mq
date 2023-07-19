import { makeAutoObservable, runInAction } from "mobx";
import authStore from "./AuthStore";
import temperStore from "./TemperStore";

class HostStore {
  // host;
  // port;
  // login;
  // password;
  // obj1;
  // obj2;
  constructor() {
    this.host = "";
    this.port = "";
    this.login = "";
    this.obj1 = "";
    this.obj2 = "";
    this.obj3 = "";
    this.password = "";
    //this.getHostFromStorage();
    makeAutoObservable(this, {});
  }

  getHostFromStorage() {
    let _host = localStorage.getItem("host") || "";
    if (_host === "") _host = "vit496.ru";
    let _port = localStorage.getItem("port") || "";
    if (_port === "") _port = "9083";
    let _login = localStorage.getItem("login") || "";
    if (_login === "") _login = "ab@m.ru";
    let _obj1 = localStorage.getItem("obj1") || "";
    if (_obj1 === "") _obj1 = "0802";
    let _obj2 = localStorage.getItem("obj2") || "";
    if (_obj2 === "") _obj2 = "0803";
    let _obj3 = localStorage.getItem("obj3") || "";
    if (_obj3 === "") _obj3 = "0804";
    let _pass = localStorage.getItem("pass") || "";

    this.host = _host;
    this.port = _port;
    this.login = _login;
    this.obj1 = _obj1;
    this.obj2 = _obj2;
    this.obj3 = _obj3;
    this.password = _pass;
    this.updAuthStore();
    temperStore.fillNobj(0, this.obj1);
    temperStore.fillNobj(1, this.obj2);
    temperStore.fillNobj(2, this.obj3);
  }

  saveHostToStorage() {
    localStorage.setItem("host", this.host);
    localStorage.setItem("port", this.port);
    localStorage.setItem("obj1", this.obj1);
    localStorage.setItem("obj2", this.obj2);
    localStorage.setItem("obj3", this.obj3);
    this.saveLoginToStorage();
    temperStore.fillNobj(0, this.obj1);
    temperStore.fillNobj(1, this.obj2);
    temperStore.fillNobj(2, this.obj3);
  }
  saveLoginToStorage() {
    localStorage.setItem("login", this.login);
    localStorage.setItem("pass", this.password);
    this.updAuthStore();
  }

  updAuthStore() {
    if (this.password !== "") {
      authStore.setIsPassPresent(true);
    } else {
      authStore.setIsPassPresent(false);
    }
  }

  setObj1(v) {
    this.obj1 = v.target.value;
    temperStore.fillNobj(0, this.obj1);
  }
  setObj2(v) {
    this.obj2 = v.target.value;
    temperStore.fillNobj(1, this.obj2);
  }
  setObj3(v) {
    this.obj3 = v.target.value;
    temperStore.fillNobj(2, this.obj3);
  }
  setHost(v) {
    this.host = v.target.value;
  }
  setPort(v) {
    this.port = v.target.value;
  }
  setLogin(v) {
    console.log("set login");
    runInAction(() => {
      this.login = v.target.value;
    });
    //this.login = v.target.value;
  }
  setPassword(v) {
    this.password = v.target.value;
  }
}

const hostStore = new HostStore();

export default hostStore;
