/* eslint-disable react-hooks/exhaustive-deps */
import { observer } from "mobx-react-lite";
import React, { useEffect, useState } from "react";
import temperStore from "../../store/TemperStore";
import OutStatus from "./OutStatus";
import AjustTemper from "./AjustTemper";
import Temper from "./Temper";

const SetOut = observer(({ indObj, indOut }) => {
  return (
    <div className="row">
      <div className="col-md-4">
        <div className="mb-2">
          <hr />
          {temperStore._nvobj[indObj].valid && (
            <>
              <OutStatus indObj={indObj} indOut={indOut} />
              <AjustTemper indObj={indObj} indOut={indOut} />
            </>
          )}
        </div>
        <Temper indObj={indObj} indOut={indOut} />
      </div>
    </div>
  );
});

export default SetOut;
