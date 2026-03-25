import { Controls, ReactFlow } from "@xyflow/react";

import "@xyflow/react/dist/style.css";
import {
  ContextMenu,
  ContextMenuContent,
  ContextMenuGroup,
  ContextMenuItem,
  ContextMenuTrigger,
} from "~/components/atoms/contextmenu/ContextMenu";

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
      <ContextMenu>
        <ContextMenuTrigger>
          <ReactFlow
            nodes={nodes}
            edges={[]}
            proOptions={{ hideAttribution: true }}
            maxZoom={0.5}
          >
            <Controls />
          </ReactFlow>
        </ContextMenuTrigger>
        <ContextMenuContent className="w-48">
          <ContextMenuGroup>
            <ContextMenuItem>Add Anchor Point</ContextMenuItem>
          </ContextMenuGroup>
        </ContextMenuContent>
      </ContextMenu>
    </div>
  );
}
