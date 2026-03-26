import React from "react";
import { type Node, type NodeProps } from "@xyflow/react";
import { cn } from "~/utils/cn";

type DeviceNode = Node<{
  id: number;
  status: "active" | "inactive";
}>;

export default function DeviceNode({ data }: NodeProps<DeviceNode>) {
  return (
    <div className="relative w-[150px] rounded-md">
      <p
        className={cn(
          "absolute top-6 right-0 rounded-md border-2 px-2 py-1",
          data.status === "inactive" &&
            "border-green-800 bg-green-200 text-green-950",
          data.status === "active" && "border-red-800 bg-red-200 text-red-950",
        )}
      >
        {data.status === "inactive" ? "Available" : "In use"}
      </p>

      <img src="/device.png" alt="device" />
    </div>
  );
}
