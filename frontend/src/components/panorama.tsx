"use client";

import React, { useState } from "react";
import { Pannellum } from "pannellum-react";

export default function PannellumReact() {
  const [currentPanorama, setCurrentPanorama] = useState("/room.jpg");

  return (
    <div>
      <p className="text-lg font-bold mb-4 mt-4">
        Главный корпус, г. Донецк, ул. Пушкина, д. 1
      </p>
      <Pannellum
        width="850px"
        height="500px"
        image={currentPanorama}
        pitch={10}
        yaw={180}
        hfov={100}
        autoLoad
      >
        <Pannellum.Hotspot
          type="custom"
          pitch={-10}
          yaw={-120}
          handleClick={(evt: Event, name: string) =>
            setCurrentPanorama(
              currentPanorama !== "/room.jpg" ? "/room.jpg" : "/mus.jpg"
            )
          }
        />
        <Pannellum.Hotspot
          type="info"
          pitch={-10}
          yaw={100}
          text="Главна библиотека ДонНТУ"
          URL="https://donntu.ru/library?ysclid=lv89zetexz657935502"
        />
      </Pannellum>
    </div>
  );
}
