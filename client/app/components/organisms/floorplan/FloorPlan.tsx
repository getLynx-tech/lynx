import React, { useMemo } from "react";
import {
  Controls,
  ReactFlow,
  useNodesState,
  useReactFlow,
  ReactFlowProvider,
  type NodeTypes,
} from "@xyflow/react";

import "@xyflow/react/dist/style.css";
import {
  ContextMenu,
  ContextMenuContent,
  ContextMenuGroup,
  ContextMenuItem,
  ContextMenuTrigger,
} from "~/components/atoms/contextmenu/ContextMenu";
import { Anchor } from "lucide-react";
import { useState, useCallback } from "react";
import BackgroundNode from "~/components/atoms/node/BackgroundNode";
import AnchorNode from "~/components/atoms/node/AnchorNode";

const initialNodes = [
  {
    id: "bg",
    type: "background",
    position: { x: 0, y: 0 },
    draggable: false,
    selectable: false,
    data: {},
  },
];

function FloorPlanInner() {
  const [nodes, setNodes, onNodesChange] = useNodesState(initialNodes);
  const [lastClickPosition, setLastClickPosition] = useState<{
    x: number;
    y: number;
  } | null>(null);
  const { screenToFlowPosition } = useReactFlow();

  const handleContextMenu = useCallback(
    (event: React.MouseEvent) => {
      const position = screenToFlowPosition({
        x: event.clientX,
        y: event.clientY,
      });
      setLastClickPosition(position);
    },
    [screenToFlowPosition],
  );

  const nodeTypes: NodeTypes = useMemo(() => {
    return {
      background: BackgroundNode,
      anchor: AnchorNode,
    };
  }, []);

  const handleAddAnchor = useCallback(() => {
    if (!lastClickPosition) return;

    setNodes((prev) => [
      ...prev,
      {
        id: `anchor-${Date.now()}`,
        type: "anchor",
        position: lastClickPosition,
        draggable: true,
        selectable: true,
        data: {},
      },
    ]);
  }, [lastClickPosition, setNodes]);

  return (
    <div className="h-full w-full bg-white">
      <ContextMenu>
        <ContextMenuTrigger asChild>
          <div className="h-full w-full" onContextMenu={handleContextMenu}>
            <ReactFlow
              nodeTypes={nodeTypes}
              nodes={nodes}
              edges={[]}
              onNodesChange={onNodesChange}
              proOptions={{ hideAttribution: true }}
              maxZoom={0.5}
            >
              <Controls />
            </ReactFlow>
          </div>
        </ContextMenuTrigger>
        <ContextMenuContent className="w-48">
          <ContextMenuGroup>
            <ContextMenuItem className="flex gap-2" onSelect={handleAddAnchor}>
              <Anchor />
              Add Anchor Point
            </ContextMenuItem>
          </ContextMenuGroup>
        </ContextMenuContent>
      </ContextMenu>
    </div>
  );
}

export default function FloorPlan() {
  return (
    <ReactFlowProvider>
      <FloorPlanInner />
    </ReactFlowProvider>
  );
}
