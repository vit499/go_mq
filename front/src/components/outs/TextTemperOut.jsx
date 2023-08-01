import { observer } from "mobx-react-lite";
import React, { useEffect, useState } from "react";
import temperStore from "../../store/TemperStore";

const TextTemperOut = observer(({ indObj, indOut }) => {
  return <div>{`${temperStore.getFtOut(indObj, indOut)}`}</div>;
});

export default TextTemperOut;
