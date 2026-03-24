import React from "react";
import type { Route } from "./+types/home";
import { auth } from "~/utils/auth/server";
import { redirect, useLoaderData } from "react-router";
import MapView from "~/components/atoms/map/MapView";
import { serverEnv } from "~/env.server";

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

  const mapBoxAccessToken = serverEnv.MAPBOX_ACCESS_TOKEN;

  return { mapBoxAccessToken };
}

export default function Home() {
  const { mapBoxAccessToken } = useLoaderData<typeof loader>();

  return (
    <div className="h-screen w-screen">
      <MapView mapBoxAccessToken={mapBoxAccessToken} />
    </div>
  );
}
