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
      <div className="d-flex flex-row">
        <button className="me-2" onClick={onMinus}>
          -
        </button>
        <div className="mt-1 ms-1">{` ${temperStore.getFtOut(
          indObj,
          indOut
        )}`}</div>
        <button className="ms-2 me-3" onClick={onPlus}>
          +
        </button>
        <button onClick={onSet}>Установить</button>
      </div>
    </div>
  );
});

export default AjustTemper;
