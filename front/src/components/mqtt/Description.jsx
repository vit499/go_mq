import { observer } from "mobx-react-lite";
import React from "react";
import descrStore from "../../store/DescrStore";

const Description = observer(() => {
  const handleSave = () => {
    descrStore.saveDescrToStorage();
  };
  return (
    <div className="container">
      <div>Описание выходов</div>
      {descrStore.outs.map((obj, indObj) => (
        <div key={indObj}>
          {obj.map((out, indOut) => (
            <div key={indOut} className="mb-2 row">
              <label className="col-sm-2">{`объект${indObj + 1} выход${
                indOut + 1
              }`}</label>
              <div className="col-sm-2">
                <input
                  onChange={(v) => descrStore.setDescrOut(indObj, indOut, v)}
                  value={descrStore.outs[indObj][indOut]}
                />
              </div>
            </div>
          ))}
        </div>
      ))}
      <hr />
      <div>Описание температур</div>
      {descrStore.tempers.map((obj, indObj) => (
        <div key={indObj}>
          {obj.map((temper, indTemper) => (
            <div key={indTemper} className="mb-2 row">
              <label className="col-sm-2">{`объект${indObj + 1} датчик${
                indTemper + 1
              }`}</label>
              <div className="col-sm-2">
                <input
                  onChange={(v) =>
                    descrStore.setDescrTemper(indObj, indTemper, v)
                  }
                  value={descrStore.tempers[indObj][indTemper]}
                />
              </div>
            </div>
          ))}
        </div>
      ))}

      <div className="mb-2 row">
        <label className="col-sm-2"></label>
        <div className="col-sm-2">
          <button onClick={handleSave}>Save</button>
        </div>
      </div>
    </div>
  );
});

export default Description;
