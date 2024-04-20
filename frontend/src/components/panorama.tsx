"use client";

import React, { useEffect, useState } from "react";

// @ts-ignore
import { Pannellum } from "pannellum-react";

export default function PannellumReact() {
  const [currentPanorama, setCurrentPanorama] = useState("/room.jpg");

  return (
    <div>
      {typeof window !== "undefined" && (
        <Pannellum
          width="1000px"
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
            name="hs1"
          />
        </Pannellum>
      )}
    </div>
  );
}
