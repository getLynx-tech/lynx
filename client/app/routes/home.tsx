import React from "react";
import type { Route } from "./+types/home";
import { auth } from "~/utils/auth/server";
import { redirect } from "react-router";
import { ReactFlowProvider } from "@xyflow/react";
import FloorPlan from "~/components/organisms/floorplan/FloorPlan";
import ScaleInput from "~/components/molecules/input/ScaleInput";
import { useTRPC } from "~/utils/trpc/react";
import { useQuery } from "@tanstack/react-query";

export function meta() {
  return [{ title: "Lynx" }];
}

export async function loader(loaderArgs: Route.LoaderArgs) {
  const { request } = loaderArgs;

  const session = await auth.api.getSession({
    headers: request.headers,
  });

  if (!session) {
    return redirect("/login");
  }
}

export default function Home() {
  const trpc = useTRPC();
  const { data: scaleData } = useQuery(trpc.scale.getScale.queryOptions());

  return (
    <div className="h-full w-full">
      <div className="fixed top-4 left-4 z-50">
        {scaleData && (
          <ScaleInput
            defaultValues={{
              meters: scaleData?.meters ?? 0,
              pixels: scaleData?.pixels ?? 0,
            }}
          />
        )}
      </div>
      <div className="fixed top-0 left-0 h-screen w-screen">
        <ReactFlowProvider>
          <FloorPlan />
        </ReactFlowProvider>
      </div>
    </div>
  );
}
