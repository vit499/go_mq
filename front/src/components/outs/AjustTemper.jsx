import { observer } from "mobx-react-lite";
import React, { useEffect, useState } from "react";
import temperStore from "../../store/TemperStore";
import TextTemperOut from "./TextTemperOut";

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
      {/* <div>{`${temperStore.getFtOut(indObj, indOut)}`}</div> */}
      <TextTemperOut indObj={indObj} indOut={indOut} />
      <button className="ms-2 me-3" onClick={() => onPlus()}>
        +
      </button>
      <button onClick={onSet}>Установить</button>
    </div>
  );
});

export default AjustTemper;
