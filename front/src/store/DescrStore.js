import { makeAutoObservable, runInAction } from "mobx";

class DescrStore {
  constructor() {
    this.outs = [
      ["", "шкаф", "", ""],
      ["", "окно", "чердак", ""],
      ["подвал", "", "", ""],
    ];
    this.tempers = [
      ["шкаф", "", ""],
      ["коридор", "чердак", ""],
      ["улица", "подвал", ""],
    ];
    //this.getDescrFromStorage();
    makeAutoObservable(this, {});
  }

  getDescrFromStorage() {
    for (let obj = 0; obj < 3; obj++) {
      for (let out = 0; out < 4; out++) {
        let item = `out${obj}${out}`;
        let d = localStorage.getItem(item) || "";
        //console.log(`get item=${item} d=${d}`);
        if (d !== "") {
          this.outs[obj][out] = d;
        }
      }
    }
    for (let obj = 0; obj < 3; obj++) {
      for (let t = 0; t < 3; t++) {
        let item = `temper${obj}${t}`;
        let d = localStorage.getItem(item) || "";
        if (d !== "") {
          this.tempers[obj][t] = d;
        }
      }
    }
    //console.log(`outs =`, JSON.stringify(this.outs, null, 2));
  }

  saveDescrToStorage() {
    for (let obj = 0; obj < 3; obj++) {
      for (let out = 0; out < 4; out++) {
        const item = `out${obj}${out}`;
        const d = this.outs[obj][out];
        localStorage.setItem(item, d);
        //console.log(`save item=${item} d=${d}`);
      }
    }
    for (let obj = 0; obj < 3; obj++) {
      for (let t = 0; t < 3; t++) {
        const item = `temper${obj}${t}`;
        const d = this.tempers[obj][t];
        localStorage.setItem(item, d);
      }
    }
  }

  setDescrOut(obj, out, e) {
    const d = e.target.value;
    runInAction(() => {
      this.outs[obj][out] = d;
    });
  }
  setDescrTemper(obj, t, e) {
    const d = e.target.value;
    runInAction(() => {
      this.tempers[obj][t] = d;
    });
  }
  getDescrOut(obj, out) {
    return this.outs[obj][out];
  }
  getDescrTemper(obj, t) {
    return this.tempers[obj][t];
  }
}

const descrStore = new DescrStore();

export default descrStore;
