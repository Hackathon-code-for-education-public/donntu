"use client";

import React from "react";

// @ts-ignore
import { Pannellum } from "pannellum-react";
import { Panorama } from "@/api/panorama";
import { S3_HOST } from "@/lib/config";

interface PanoramaProps {
  panorama: Panorama;
}

export default function PannellumReact({ panorama }: PanoramaProps) {
  return (
    <div>
      <p className="text-lg font-bold mb-4 mt-4">
        {panorama.name + ", " + panorama.address}
      </p>
      <Pannellum
        width="850px"
        height="500px"
        image={`${S3_HOST}/${panorama.firstLocation}`}
        pitch={10}
        yaw={180}
        hfov={100}
        autoLoad
      ></Pannellum>
    </div>
  );
}
