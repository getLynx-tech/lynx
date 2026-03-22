import {
  type RouteConfig,
  index,
  layout,
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
  layout("routes/auth/layout.tsx", [
    route("signup", "routes/auth/signup.tsx"),
    route("login", "routes/auth/login.tsx"),
  ]),
] satisfies RouteConfig;
