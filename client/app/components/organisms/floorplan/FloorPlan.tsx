import React, { useEffect, useMemo } from "react";
import {
  Controls,
  ReactFlow,
  useNodesState,
  useReactFlow,
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
import { useTRPC } from "~/utils/trpc/react";
import { useMutation, useQuery } from "@tanstack/react-query";
import { Spinner } from "~/components/ui/spinner";
import DeviceNode from "~/components/atoms/node/DeviceNode";

const backgroundNode = {
  id: "bg",
  type: "background",
  position: { x: 0, y: 0 },
  draggable: false,
  selectable: false,
  data: {},
};

export default function FloorPlan() {
  const trpc = useTRPC();

  const { data: anchorNodesData } = useQuery(
    trpc.anchorRouter.getAllAnchors.queryOptions(),
  );

  const { data: deviceNodesData } = useQuery(
    trpc.deviceRouter.getAllDevices.queryOptions(undefined, {
      refetchInterval: 500,
    }),
  );

  const upsertAnchorNodes = useMutation(
    trpc.anchorRouter.upsertAnchors.mutationOptions(),
  );

  const [nodes, setNodes, onNodesChange] = useNodesState([backgroundNode]);
  const [lastClickPosition, setLastClickPosition] = useState<{
    x: number;
    y: number;
  } | null>(null);
  const { screenToFlowPosition } = useReactFlow();

  useEffect(() => {
    const anchorNodes = anchorNodesData?.map((node) => ({
      id: node.id,
      type: "anchor",
      position: { x: node.x, y: node.y },
      draggable: true,
      selectable: true,
      data: {
        id: anchorNodesData.find((n) => n.id === node.id)?.id || "",
      },
    }));
    const deviceNodes = deviceNodesData?.map((node) => ({
      id: node.id,
      type: "device",
      position: { x: node.x, y: node.y },
      draggable: true,
      selectable: true,
      data: {
        status: node.status === "available" ? "available" : "in use",
      },
    }));
    setNodes([backgroundNode, ...(deviceNodes ?? []), ...(anchorNodes ?? [])]);
  }, [anchorNodesData, deviceNodesData]);

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
      device: DeviceNode,
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

  const onSubmit = useCallback(() => {
    const anchors = nodes
      .filter((node) => node.type === "anchor")
      .map((node) => ({
        id: node.id,
        x: node.position.x,
        y: node.position.y,
      }));

    upsertAnchorNodes.mutate(anchors);
  }, [nodes, upsertAnchorNodes]);

  return (
    <>
      <div className="fixed top-4 right-4 z-50">
        <button
          onClick={onSubmit}
          disabled={upsertAnchorNodes.isPending}
          className="flex cursor-pointer items-center gap-2 rounded-md bg-black px-2 py-1 text-sm text-white"
        >
          Save Floorplan
          {upsertAnchorNodes.isPending && <Spinner />}
        </button>
      </div>
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
              <ContextMenuItem
                className="flex gap-2"
                onSelect={handleAddAnchor}
              >
                <Anchor />
                Add Anchor Point
              </ContextMenuItem>
            </ContextMenuGroup>
          </ContextMenuContent>
        </ContextMenu>
      </div>
    </>
  );
}
