import React from "react";
import { observer } from "mobx-react-lite";
import hostStore from "../../store/HostStore";

const Settings = observer(() => {
  const handleSave = () => {
    hostStore.saveHostToStorage();
  };
  return (
    <div className="container">
      <div className="mb-2 row">
        <label className="col-sm-2">Host</label>
        <div className="col-sm-2">
          <input
            onChange={(v) => hostStore.setHost(v)}
            value={hostStore.host}
          />
        </div>
      </div>
      <div className="mb-2 row">
        <label className="col-sm-2">Port</label>
        <div className="col-sm-2">
          <input
            onChange={(v) => hostStore.setPort(v)}
            value={hostStore.port}
          />
        </div>
      </div>

      <div className="mb-2 row">
        <label className="col-sm-2">Object</label>
        <div className="col-sm-2">
          <input
            onChange={(v) => hostStore.setObj1(v)}
            value={hostStore.obj1}
          />
        </div>
      </div>
      <div className="mb-2 row">
        <label className="col-sm-2">Object</label>
        <div className="col-sm-2">
          <input
            onChange={(v) => hostStore.setObj2(v)}
            value={hostStore.obj2}
          />
        </div>
      </div>
      <div className="mb-2 row">
        <label className="col-sm-2">Object</label>
        <div className="col-sm-2">
          <input
            onChange={(v) => hostStore.setObj3(v)}
            value={hostStore.obj3}
          />
        </div>
      </div>
      <div className="mb-2 row">
        <label className="col-sm-2">Login</label>
        <div className="col-sm-2">
          <input
            onChange={(v) => hostStore.setLogin(v)}
            value={hostStore.login}
          />
        </div>
      </div>
      <div className="mb-2 row">
        <label className="col-sm-2">Password</label>
        <div className="col-sm-2">
          <input
            type="password"
            onChange={(v) => hostStore.setPassword(v)}
            value={hostStore.password}
          />
        </div>
      </div>
      <div className="mb-2 row">
        <label className="col-sm-2"></label>
        <div className="col-sm-2">
          <button onClick={handleSave}>Save</button>
        </div>
      </div>
    </div>
  );
});

export default Settings;
