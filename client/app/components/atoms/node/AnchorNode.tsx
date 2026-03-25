import React from "react";
import { type Node } from "@xyflow/react";
import { Anchor } from "lucide-react";

type AnchorNode = Node<{}>;

export default function AnchorNode() {
  return (
    <div className="flex h-[80px] w-[80px] items-center justify-center rounded-md bg-blue-600 text-white shadow-xl">
      <Anchor className="h-12 w-12" />
    </div>
  );
}
