import { observer } from "mobx-react-lite";
import React, { useEffect, useState } from "react";
import mq from "../../store/Mq";
import temperStore from "../../store/TemperStore";
import TemperCompModern from "./TemperCompModern";

const SetOutModern = observer(({ indObj, indOut }) => {
  const [temperOn, setTemperOn] = useState(0);
  const [temperOnInit, setTemperOnInit] = useState(false); //
  const onPlus = () => {
    setTemperOn((v) => v + 1);
  };
  const onMinus = () => {
    setTemperOn((v) => v - 1);
  };
  const onSet = () => {
    const mes = `setout${indOut + 1}=${temperOn.toString()}`;
    mq.mqttPublish({ indObj: indObj, payload: mes });
  };
  useEffect(() => {
    // для начельной инициализации температуры включения
    // console.log("useEff temper valid", temperStore._nvobj[indObj].valid, temperOnInit);
    if (temperStore._nvobj[indObj].valid) {
      if (!temperOnInit) {
        // устанавливается один раз при первом получении данных
        // console.log("useEff temper valid", temperStore._nvobj[indObj].valid, temperStore._nvobj[indObj].ftout[indOut]);
        setTemperOn(temperStore._nvobj[indObj].ftout[indOut]);
        setTemperOnInit(true);
      }
    }
  }, [temperStore._nvobj[indObj].valid]);
  return (
    <div className="">
      <div className="">
        <div className="mb-2">
          {mq.connectStatus === "Connected" &&
            !temperStore._nvobj[indObj].valid && <Loading />}
          {temperStore._nvobj[indObj].valid && (
            <>
              <hr />
              <div className="mb-2">
                {temperStore._nvobj[indObj].sout[indOut] !== 0 ? (
                  <div className="vihod">
                    <span className="marg down-line">
                      <ion-icon name="snow-outline"></ion-icon>
                    </span>
                    <span className="vihod-vk2 down-line">
                      <ion-icon name="sunny-outline"></ion-icon>
                    </span>
                  </div>
                ) : (
                  <div className="vihod">
                    <span className="vihod-vk1 down-line">
                      <ion-icon name="snow-outline"></ion-icon>
                    </span>
                    <span className="marg down-line">
                      <ion-icon name="sunny-outline"></ion-icon>
                    </span>
                  </div>
                )}
              </div>
              <div className="mb-2 up-T up-T-mini">
                {" "}
                Включение при{" "}
                <ion-icon
                  id="qwe"
                  name="thermometer-outline"
                ></ion-icon> ниже{" "}
                <span className="text-yellow ">
                  {temperStore._nvobj[indObj].ftout[indOut]}
                </span>
                °C
              </div>
              {/* --pult-- */}
              <div className="pult">
                <div className="mb-1 p-2 butter">
                  <div className="l1">
                    <button type="button" className="but-min" onClick={onMinus}>
                      <ion-icon name="remove-circle-outline"></ion-icon>
                    </button>
                    <span className="text-yellow">{temperOn.toString()}</span>
                    <button type="button" className="but-pl" onClick={onPlus}>
                      <ion-icon name="add-circle-outline"></ion-icon>
                    </button>
                  </div>
                  <button type="button" className="click " onClick={onSet}>
                    Установить
                  </button>
                </div>
                <TemperCompModern indObj={indObj} indOut={indOut} />
              </div>
              {/* --pult-- */}
            </>
          )}
        </div>
      </div>
    </div>
  );
});

export default SetOutModern;
