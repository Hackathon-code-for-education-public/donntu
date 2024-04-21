"use client";

import React, { useState } from "react";

// @ts-ignore
import { Pannellum } from "pannellum-react";
import { Panorama } from "@/api/panorama";

interface PanoramaProps {
  panorama: Panorama;
}

export default function PannellumReact({ panorama }: PanoramaProps) {
  const [currentPanorama, setCurrentPanorama] = useState(panorama.firstLocation);

  return (
    <div>
      <p className="text-lg font-bold mb-4 mt-4">
        {panorama.name + ", " + panorama.address}
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
              currentPanorama !== panorama.firstLocation ? panorama.firstLocation : panorama.secondLocation
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
