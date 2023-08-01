/* eslint-disable react-hooks/exhaustive-deps */
import { observer } from "mobx-react-lite";
import React, { useEffect, useState } from "react";
import descrStore from "../../store/DescrStore";
//import mq from "../../store/Mq";
import temperStore from "../../store/TemperStore";
//import wsStore from "../../store/WsStore";

const SetOut = observer(({ indObj, indOut }) => {
  //const [temperOn, setTemperOn] = useState(0);
  const [temperOnInit, setTemperOnInit] = useState(false); //
  const onPlus = () => {
    //setTemperOn((v) => v + 1);
    temperStore.plusFtoutCopy(indObj, indOut);
  };
  const onMinus = () => {
    //setTemperOn((v) => v - 1);
    temperStore.minusFtoutCopy(indObj, indOut);
  };
  const onSet = () => {
    // const mes = `setout${indOut + 1}=${temperOn.toString()}`;
    // wsStore.WsPublish({ indObj: indObj, payload: mes });
    temperStore.SetFtout(indObj, indOut);
  };
  // useEffect(() => {
  //   // для начельной инициализации температуры включения
  //   console.log(
  //     `set temper on, ind=${indObj}  valid=${temperStore._nvobj[indObj].valid}, t=${temperStore._nvobj[indObj].ftout[indOut]}`
  //   );
  //   if (temperStore._nvobj[indObj].valid) {
  //     if (!temperOnInit) {
  //       // устанавливается один раз при первом получении данных
  //       console.log(
  //         `set temper on, ind=${indObj}  valid=${temperStore._nvobj[indObj].valid}, t=${temperStore._nvobj[indObj].ftout[indOut]}`
  //       );
  //       setTemperOn(temperStore._nvobj[indObj].ftout[indOut]);
  //       setTemperOnInit(true);
  //     }
  //   }
  // }, [temperStore._nvobj[indObj].valid]);
  return (
    <div className="row">
      <div className="col-md-4">
        <div className="mb-2">
          <hr />
          {temperStore._nvobj[indObj].valid && (
            <>
              <div className="mb-2">
                {!temperStore._nvobj[indObj].online && (
                  <div style={{ backgroundColor: "#dddddd" }}>
                    {`не на связи`}
                  </div>
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
              </div>
              <div className="mb-2">{` включение при Т ниже ${temperStore._nvobj[indObj].ftout[indOut]} `}</div>
              <div className="mb-2">
                <button className="me-2" onClick={onMinus}>
                  -
                </button>
                {/* {` ${temperOn.toString()}`} */}
                {` ${temperStore._nvobj[indObj].ftout_copy[indOut]}`}
                <button className="ms-2 me-3" onClick={onPlus}>
                  +
                </button>
                <button onClick={onSet}>Установить</button>
              </div>
            </>
          )}
        </div>
        <div>{` Температура=${temperStore.getTemper(
          temperStore._nvobj[indObj].ind,
          indOut
        )} `}</div>
      </div>
    </div>
  );
});

export default SetOut;
