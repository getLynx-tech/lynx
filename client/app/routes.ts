import {
  type RouteConfig,
  index,
  prefix,
  route,
} from "@react-router/dev/routes";

export default [
  index("routes/home.tsx"),
  ...prefix("api", [
    route("auth/*", "routes/api/auth.ts"),

    // tRPC routes
    route("trpc/*", "routes/api/trpc.ts"),
  ]),
] satisfies RouteConfig;
