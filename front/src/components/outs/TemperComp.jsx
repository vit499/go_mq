import { observer } from "mobx-react-lite";
import React from "react";
import descrStore from "../../store/DescrStore";
import temperStore from "../../store/TemperStore";

const TemperComp = observer(() => {
  return (
    <div className="container">
      <div className="row mb-2">
        <div className="col-md-4">
          <hr />
          {temperStore._nvobj.map((o) => (
            <div key={o.ind}>
              {o.temper.map((t, ind) => (
                <div key={ind}>
                  {t !== 0x80 && (
                    // <div>{`T=${t} (object=${o.nobj} sensor=${ind + 1}) `}</div>
                    <div>{`T=${t} (${descrStore.tempers[o.ind][ind]}) `}</div>
                  )}
                </div>
              ))}
            </div>
          ))}
        </div>
      </div>
    </div>
  );
});

export default TemperComp;
