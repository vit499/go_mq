import { makeAutoObservable, runInAction } from "mobx";
import devSend from "../utils/devsend";
import wsStore from "../store/WsStore";

// const nvobj = {
//   fout: [0, 0, 0, 0],
//   ftout: [0, 0, 0, 0],
//   sout: [0, 0, 0, 0],
//   temper: [0x80, 0x80, 0x80],
// };
class TemperStore {
  constructor() {
    this._nvobj = [
      {
        ind: 0,
        nobj: "0808",
        fout: [0, 0, 0, 0],
        ftout: [0, 0, 0, 0],
        sout: [0, 0, 0, 0],
        indtemp: [0, 0, 0, 0],
        temper: [0x80, 0x80, 0x80],
        online: false,
        valid: false,
        ftout_copy: [0, 0, 0, 0],
      },
      {
        ind: 1,
        nobj: "0809",
        fout: [0, 0, 0, 0],
        ftout: [0, 0, 0, 0],
        sout: [0, 0, 0, 0],
        indtemp: [0, 0, 0, 0],
        temper: [0x80, 0x80, 0x80],
        online: false,
        valid: false,
        ftout_copy: [0, 0, 0, 0],
      },
      {
        ind: 2,
        nobj: "0801",
        fout: [0, 0, 0, 0],
        ftout: [0, 0, 0, 0],
        sout: [0, 0, 0, 0],
        indtemp: [0, 0, 0, 0],
        temper: [0x80, 0x80, 0x80],
        online: false,
        valid: false,
        ftout_copy: [0, 0, 0, 0],
      },
    ];
    makeAutoObservable(this, {});
  }

  plusFtoutCopy(indObj, indOut) {
    let t = this._nvobj[indObj].ftout_copy[indOut];
    t = t + 1;
    runInAction(() => {
      //console.log(`plusFtoutCopy ${t}`);
      this._nvobj[indObj].ftout_copy[indOut] = t;
    });
  }
  minusFtoutCopy(indObj, indOut) {
    let t = this._nvobj[indObj].ftout_copy[indOut];
    t = t - 1;
    runInAction(() => {
      //console.log(`minusFtoutCopy ${t}`);
      this._nvobj[indObj].ftout_copy[indOut] = t;
    });
  }
  getFtOut(indObj, indOut) {
    const t = this._nvobj[indObj].ftout_copy[indOut];
    return t;
  }
  SetFtout(indObj, indOut) {
    let t = this._nvobj[indObj].ftout_copy[indOut];
    const mes = `setout${indOut + 1}=${t}`;
    wsStore.WsPublish({ indObj: indObj, payload: mes });
  }
  fillNobj(ind, nobj) {
    if (ind < 3) this._nvobj[ind].nobj = nobj;
    devSend.fillNobj(ind, nobj);
  }

  clear(indObj) {
    devSend.clear(indObj);
    runInAction(() => {
      this._nvobj[indObj].fout = [0, 0, 0, 0];
      this._nvobj[indObj].ftout = [0, 0, 0, 0];
      this._nvobj[indObj].sout = [0, 0, 0, 0];
      this._nvobj[indObj].indtemp = [0, 0, 0, 0];
      this._nvobj[indObj].temper = [0x80, 0x80, 0x80];
      this._nvobj[indObj].online = false;
      this._nvobj[indObj].valid = false;
      this._nvobj[indObj].ftout_copy = [0, 0, 0, 0];
    });
  }
  clearAll() {
    devSend.clearAll();
    runInAction(() => {
      this._nvobj.forEach((o) => {
        o.fout = [0, 0, 0, 0];
        o.ftout = [0, 0, 0, 0];
        o.sout = [0, 0, 0, 0];
        o.indtemp = [0, 0, 0, 0];
        o.temper = [0x80, 0x80, 0x80];
        o.online = false;
        o.valid = false;
        o.ftout_copy = [0, 0, 0, 0];
      });
    });
  }
  getTemper(indObj, indOut) {
    const indTemper = this._nvobj[indObj].indtemp[indOut];
    let t = this._nvobj[indObj].temper[indTemper] & 0xff;
    if (t === 0x80) {
      return "--";
    }
    if ((t & (1 << 7)) != 0) t = t - 256;
    return `${t} (датчик ${indTemper + 1})`;
  }
  getTemperAll() {
    let strRes = "";
    this._nvobj.forEach((o) => {
      o.temper.forEach((val, ind) => {
        let t = val;
        if (t !== 0x80) {
          if ((t & (1 << 7)) != 0) t = t - 256;
          //strRes.concat(`<br/>${t} (obj${o.ind + 1} sensor${ind + 1})`);
          strRes = strRes + `\r\n ${t} (obj${o.ind + 1} sensor${ind + 1})`;
        }
      });
    });
    //console.log("strTemp", strRes);
    return strRes;
  }

  cpyObj(obj) {
    let ind;
    if (this._nvobj[0].nobj === obj.nobj) ind = 0;
    else if (this._nvobj[1].nobj === obj.nobj) ind = 1;
    else if (this._nvobj[2].nobj === obj.nobj) ind = 2;

    runInAction(() => {
      if (obj.fout.length !== 0)
        obj.fout.forEach((f, i) => {
          //console.log(`${o.nobj} fout${i + 1}=${o.fout[i].toString()}`);
          this._nvobj[ind].fout[i] = obj.fout[i];
        });
      if (obj.ftout.length !== 0)
        obj.ftout.forEach((f, i) => {
          //console.log(`${o.nobj} ftout${i + 1}=${o.ftout[i].toString()}`);
          this._nvobj[ind].ftout[i] = obj.ftout[i];
          this._nvobj[ind].ftout_copy[i] = obj.ftout[i];
        });
      if (obj.sout.length !== 0)
        obj.sout.forEach((f, i) => {
          //console.log(`${o.nobj} sout${i + 1}=${o.sout[i].toString()}`);
          this._nvobj[ind].sout[i] = obj.sout[i];
        });
      if (obj.indtemp.length !== 0)
        obj.indtemp.forEach((f, i) => {
          //console.log(`${o.nobj} indtemp${i + 1}=${o.indtemp[i].toString()}`);
          this._nvobj[ind].indtemp[i] = obj.indtemp[i];
        });
      obj.temper.forEach((f, i) => {
        if (i < 3) {
          let t = obj.temper[i] & 0xff;
          if (t != 0x80) {
            if ((t & (1 << 7)) != 0) t = t - 256;
          }
          //console.log(`${o.nobj} temper${i + 1}=${o.temper[i].toString()}`);
          this._nvobj[ind].temper[i] = t;
        }
      });
      this._nvobj[ind].online = obj.online;
      this._nvobj[ind].valid = obj.valid;
      //console.log(`online${ind}=${this._nvobj[ind].online}`);
      //console.log(`valid${ind}=${this._nvobj[ind].valid}`);
    });
  }
  recMes(topic, message) {
    const xx = devSend.parseMes(topic, message);
    if (xx && xx.valid) {
      this.cpyObj(xx);
    }
  }

  recMesJson(topic, message) {
    const xx = devSend.parseMesJson(topic, message);
    if (xx && xx.valid) {
      this.cpyObj(xx);
    }
  }
}

const temperStore = new TemperStore();

export default temperStore;
