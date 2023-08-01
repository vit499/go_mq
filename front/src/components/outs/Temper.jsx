/* eslint-disable react-hooks/exhaustive-deps */
import { observer } from "mobx-react-lite";
import React, { useEffect, useState } from "react";
import temperStore from "../../store/TemperStore";

const Temper = observer(({ indObj, indOut }) => {
  return (
    <div>{` Температура=${temperStore.getTemper(
      temperStore._nvobj[indObj].ind,
      indOut
    )} `}</div>
  );
});

export default Temper;
