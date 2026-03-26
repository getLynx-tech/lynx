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
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "~/components/ui/dialog";
import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";

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

  const [showDialog, setShowDialog] = useState<boolean>(false);
  const [anchorId, setAnchorId] = useState("");

  const [nodes, setNodes, onNodesChange] = useNodesState([backgroundNode]);
  const [lastClickPosition, setLastClickPosition] = useState<{
    x: number;
    y: number;
  } | null>(null);
  const { screenToFlowPosition } = useReactFlow();

  useEffect(() => {
    if (!anchorNodesData) return;
    const anchorNodes = anchorNodesData.map((node) => ({
      id: node.id,
      type: "anchor",
      position: { x: node.x, y: node.y },
      draggable: true,
      selectable: true,
      data: { id: node.id },
    }));
    setNodes((prev) => {
      const nonAnchorNodes = prev.filter((n) => n.type !== "anchor");
      return [...nonAnchorNodes, ...anchorNodes];
    });
  }, [anchorNodesData]);

  useEffect(() => {
    if (!deviceNodesData) return;
    const deviceNodes = deviceNodesData.map((node) => ({
      id: node.id,
      type: "device",
      position: { x: node.x, y: node.y },
      draggable: false,
      selectable: false,
      data: {
        status: node.status,
      },
    }));
    setNodes((prev) => {
      const nonDeviceNodes = prev.filter((n) => n.type !== "device");
      return [...nonDeviceNodes, ...deviceNodes];
    });
  }, [deviceNodesData]);

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
        id: anchorId,
        type: "anchor",
        position: lastClickPosition,
        draggable: true,
        selectable: true,
        data: {
          id: anchorId,
        },
      },
    ]);

    setShowDialog(false);
  }, [lastClickPosition, setNodes, anchorId]);

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
      <Dialog open={showDialog} onOpenChange={setShowDialog}>
        <DialogContent className="sm:max-w-sm">
          <DialogHeader>
            <DialogTitle>Anchor ID</DialogTitle>
            <DialogDescription>
              Please enter the ID of the anchor point.
            </DialogDescription>
          </DialogHeader>
          <form>
            <Input
              placeholder="Anchor ID"
              value={anchorId}
              onChange={(e) => setAnchorId(e.target.value)}
            />
          </form>
          <DialogFooter>
            <DialogClose asChild>
              <Button variant="outline">Cancel</Button>
            </DialogClose>
            <Button onClick={handleAddAnchor}>Save changes</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
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
                minZoom={0.1}
                maxZoom={1}
                fitView
              >
                <Controls />
              </ReactFlow>
            </div>
          </ContextMenuTrigger>
          <ContextMenuContent className="w-48">
            <ContextMenuGroup>
              <ContextMenuItem
                className="flex gap-2"
                onSelect={() => setShowDialog(true)}
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
