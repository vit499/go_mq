import { observer } from "mobx-react-lite";
import React, { useEffect, useState } from "react";
import descrStore from "../../store/DescrStore";
import temperStore from "../../store/TemperStore";

const OutStatus = observer(({ indObj, indOut }) => {
  return (
    <div className="mb-2">
      {!temperStore._nvobj[indObj].online && (
        <div style={{ backgroundColor: "#dddddd" }}>{`не на связи`}</div>
      )}
      {temperStore._nvobj[indObj].sout[indOut] !== 0 ? (
        <div
          style={{ backgroundColor: "pink" }}
        >{`обогрев ${descrStore.outs[indObj][indOut]} включен`}</div>
      ) : (
        <div style={{ backgroundColor: "#dddddd" }}>
          {`обогрев ${descrStore.outs[indObj][indOut]} выключен`}
        </div>
      )}
      <div className="mb-2">{` включение при Т ниже ${temperStore._nvobj[indObj].ftout[indOut]} `}</div>
    </div>
  );
});

export default OutStatus;
