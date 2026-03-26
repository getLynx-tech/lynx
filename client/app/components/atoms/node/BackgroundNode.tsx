import React from "react";

export default function BackgroundNode() {
  return (
    <div
      style={{
        backgroundImage: "url('/floorplan.jpg')",
        backgroundSize: "contain",
        backgroundRepeat: "no-repeat",
      }}
      className="h-[5000px] w-[5000px]"
    ></div>
  );
}
