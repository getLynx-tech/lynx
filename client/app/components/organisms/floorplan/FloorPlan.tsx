import { Controls, ReactFlow } from "@xyflow/react";

import "@xyflow/react/dist/style.css";

export default function FloorPlan() {
  const nodes = [
    {
      id: "bg",
      position: { x: 0, y: 0 },
      draggable: false,
      selectable: false,
      data: {},
      style: {
        width: 5000,
        height: 5000,
        backgroundImage: "url('/floorplan.png')",
        backgroundSize: "contain",
        backgroundRepeat: "no-repeat",
      },
    },
  ];

  return (
    <div className="h-full w-full bg-white">
      <ReactFlow
        nodes={nodes}
        edges={[]}
        proOptions={{ hideAttribution: true }}
        maxZoom={0.5}
      >
        <Controls />
      </ReactFlow>
    </div>
  );
}
