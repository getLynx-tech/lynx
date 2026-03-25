import React from "react";
import type { Route } from "./+types/home";
import { auth } from "~/utils/auth/server";
import { redirect } from "react-router";

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
  return <div className="h-screen w-screen"></div>;
}
