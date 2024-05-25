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
        <button className=" btn btn-info" onClick={onMinus}>
          -
        </button>
        <div className="pt-1 px-3 bg-warning text-dark">{` ${temperStore.getFtOut(
          indObj,
          indOut
        )}`}</div>
        <button className=" me-3 btn btn-info" onClick={onPlus}>
          +
        </button>
        <button className="btn btn-success" onClick={onSet}>
          Установить
        </button>
      </div>
    </div>
  );
});

export default AjustTemper;
