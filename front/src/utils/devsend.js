class DevSend {
  constructor() {
    // this.dv12v = "0";
    // this.fout = [];
    // this.ftout = [];
    // this.temper = [];
    // this.sout = [];

    this._nvobj = [
      {
        nobj: "0808",
        fout: [],
        ftout: [],
        sout: [],
        indtemp: [],
        temper: [0x80, 0x80, 0x80],
        online: false,
        valid: false,
      },
      {
        nobj: "0809",
        fout: [],
        ftout: [],
        sout: [],
        indtemp: [],
        temper: [0x80, 0x80, 0x80],
        online: false,
        valid: false,
      },
      {
        nobj: "0801",
        fout: [],
        ftout: [],
        sout: [],
        indtemp: [],
        temper: [0x80, 0x80, 0x80],
        online: false,
        valid: false,
      },
    ];
  }

  fillNobj(ind, nobj) {
    if (ind < 3) this._nvobj[ind].nobj = nobj;
  }

  clear(indObj) {
    this._nvobj[indObj].fout = [];
    this._nvobj[indObj].ftout = [];
    this._nvobj[indObj].sout = [];
    this._nvobj[indObj].indtemp = [];
    this._nvobj[indObj].temper = [0x80, 0x80, 0x80];
    this._nvobj[indObj].online = false;
    this._nvobj[indObj].valid = false;
  }
  clearAll() {
    this._nvobj.forEach((o) => {
      o.fout = [];
      o.ftout = [];
      o.sout = [];
      o.indtemp = [];
      o.temper = [0x80, 0x80, 0x80];
      o.online = false;
      o.valid = false;
    });
  }

  getValFromHex(mes) {
    let val = [];

    const len = mes.length / 2;
    for (let i = 0; i < len; i++) {
      const c = "0x" + mes.substring(i * 2, i * 2 + 2);
      const b = parseInt(c, 16);
      val.push(b);
    }
    //console.log("val", JSON.stringify(val, null, 2));
    return val;
  }
  getValFromBits(mes) {
    let val = [];
    let bytes = this.getValFromHex(mes);
    bytes.forEach((b) => {
      for (let i = 0; i < 8; i++) {
        let c = 0;
        if ((b & (1 << i)) !== 0) c = 1;
        val.push(c);
      }
    });
    return val;
  }

  parseMes(topic, mes) {
    const i0 = topic.indexOf("/") + 1; // ab@m.ru/0802/devsend/xxx
    const nobj = topic.substring(i0, i0 + 4);
    let obj = this._nvobj.find((o) => o.nobj === nobj);
    // console.log(`obj from topic=${nobj}`);
    if (!obj) return null;
    const i1 = topic.indexOf("devsend/") + 8;
    const t = topic.substring(i1);
    // console.log("topic=", t);

    let arrStr = [];

    if (t === "config/fout") {
      //this.fout = this.getValFromHex(mes);
      const fout = this.getValFromHex(mes);
      fout.forEach((v, i) => {
        if (i < 4) obj.fout[i] = v;
      });
      // console.log("fout", JSON.stringify(this.fout, null, 2));
    } else if (t === "config/ftout") {
      const ftout = this.getValFromHex(mes);
      ftout.forEach((v, i) => {
        if (i < 4) obj.ftout[i] = v;
      });
      // console.log("ftout", JSON.stringify(this.ftout, null, 2));
    } else if (t === "config/indtemper") {
      const indtemp = this.getValFromHex(mes);
      indtemp.forEach((v, i) => {
        if (i < 4) obj.indtemp[i] = v;
      });
      // console.log("indtemp", JSON.stringify(this.indtemp, null, 2));
    } else if (t === "status/sout") {
      const sout = this.getValFromBits(mes);
      sout.forEach((v, i) => {
        if (i < 4) obj.sout[i] = v;
      });
      // console.log("sout", JSON.stringify(this.sout, null, 2));
    } else if (t === "status/param") {
      arrStr = mes.split("&");
      arrStr.forEach((a) => {
        const b = a.split("=");
        if (b[0] === "dv_12v") {
          this.dv12v = b[1];
        } else if (b[0] === "temper") {
          //this.temper = this.getValFromHex(b[1]);
          const temper = this.getValFromHex(b[1]);
          temper.forEach((v, i) => {
            if (i < 3) obj.temper[i] = v;
          });
          //obj.temper = temper;
        }
      });
      // console.log("arrStr", JSON.stringify(arrStr, null, 2));
    }

    obj.online = true;
    let valid = true;
    if (obj.fout.length === 0) valid = false;
    if (obj.ftout.length === 0) valid = false;
    if (obj.sout.length === 0) valid = false;
    obj.valid = valid;
    //this.log1(obj);
    return obj;
  }

  log1(obj) {
    //console.log(`${obj.nobj} =`, JSON.stringify(obj, null, 2));
  }

  /*
src={
  "StrUnit": "0803",
  "Fout": [
    0,
    0,
    0,
    0
  ],
  "Sout": [
    0,
    0,
    0,
    0
  ],
  "Ftout": [
    0,
    0,
    0,
    0
  ],
  "IndTemper": [
    0,
    0,
    0,
    0
  ],
  "Temper": [
    128,
    128,
    128
  ],
  "U12v": "-",
  "LevelGsm": 0,
  "LevelWifi": 0
}
  */
  parseMesJson(topic, mes) {
    const nobj = topic;
    let obj = this._nvobj.find((o) => o.nobj === nobj);
    // console.log(`obj from topic=${nobj}`);
    if (!obj) return null;
    const src = JSON.parse(mes);

    //console.log(`src=${JSON.stringify(src, null, 2)}`);

    const Fout = src.Fout;
    if (Fout != null) {
      Fout.forEach((v, i) => {
        if (i < 4) obj.fout[i] = v;
      });
    }

    const Sout = src.Sout;
    //console.log(`Sout=${JSON.stringify(Sout, null, 2)}`);
    if (Sout != null) {
      Sout.forEach((v, i) => {
        if (i < 4) obj.sout[i] = v;
      });
    }

    const Ftout = src.Ftout;
    if (Ftout != null) {
      Ftout.forEach((v, i) => {
        if (i < 4) obj.ftout[i] = v;
      });
    }

    const IndTemper = src.IndTemper;
    if (IndTemper != null) {
      IndTemper.forEach((v, i) => {
        if (i < 4) obj.indtemp[i] = v;
      });
    }

    const Temper = src.Temper;
    if (Temper != null) {
      Temper.forEach((v, i) => {
        if (i < 3) obj.temper[i] = v;
      });
    }

    // this.dv12v = src.U12v;
    // this.LevelGsm = src.LevelGsm;
    // this.LevelWifi = src.LevelWifi;
    // this.Online = src.Online;
    obj.dv12v = src.U12v;
    obj.LevelGsm = src.LevelGsm;
    obj.LevelWifi = src.LevelWifi;
    obj.online = src.Online;

    let valid = true;
    if (obj.fout.length === 0) valid = false;
    if (obj.ftout.length === 0) valid = false;
    if (obj.sout.length === 0) valid = false;
    obj.valid = valid;
    this.log1(obj);
    return obj;
  }
}

const devSend = new DevSend();

export default devSend;
