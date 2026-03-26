import React from "react";
import { type Node, type NodeProps } from "@xyflow/react";
import { Anchor } from "lucide-react";

type AnchorNode = Node<{
  id: string;
}>;

export default function AnchorNode({ data }: NodeProps<AnchorNode>) {
  return (
    <div className="relative -translate-x-1/2 -translate-y-1/2">
      <div className="flex h-[50px] w-[50px] items-center justify-center rounded-md bg-blue-600 text-white shadow-xl">
        <Anchor className="h-10 w-10" />
      </div>

      <p className="absolute top-full left-1/2 mt-1 -translate-x-1/2 rounded bg-white px-2 py-1 shadow">
        {data.id}
      </p>
    </div>
  );
}
