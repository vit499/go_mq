import { makeAutoObservable } from "mobx";

class AuthStore {
  isAuth;
  isPassPresent;
  constructor() {
    this.isAuth = false;
    this.isPassPresent = false;
    makeAutoObservable(this, {});
  }

  // setAuth(v) {
  //   runInAction(() => {
  //     this.isAuth = v;
  //   });
  // }
  setIsPassPresent(v) {
    this.isPassPresent = v;
    this.isAuth = v;
  }
}

const authStore = new AuthStore();

export default authStore;
