import { useTRPC } from "~/utils/trpc/react";
import type { Route } from "./+types/home";
import { useQuery } from "@tanstack/react-query";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "New React Router App" },
    { name: "description", content: "Welcome to React Router!" },
  ];
}

export default function Home() {
  const trpc = useTRPC();
  const { data: rootData } = useQuery(trpc.root.getRoot.queryOptions());

  return <p>{rootData?.message}</p>;
}
