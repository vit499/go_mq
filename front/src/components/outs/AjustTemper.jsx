import { observer } from "mobx-react-lite";
import React, { useEffect, useState } from "react";
import temperStore from "../../store/TemperStore";

const AjustTemper = observer(({ indObj, indOut }) => {
  const onPlus = () => {
    temperStore.plusFtoutCopy(indObj, indOut);
  };
  const onMinus = () => {
    temperStore.minusFtoutCopy(indObj, indOut);
  };
  const onSet = () => {
    temperStore.SetFtout(indObj, indOut);
  };
  return (
    <div className="mb-2">
      <button className="me-2" onClick={() => onMinus()}>
        -
      </button>
      {` ${temperStore._nvobj[indObj].ftout_copy[indOut]}`}
      <button className="ms-2 me-3" onClick={() => onPlus()}>
        +
      </button>
      <button onClick={onSet}>Установить</button>
    </div>
  );
});

export default AjustTemper;
