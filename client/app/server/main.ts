import { createTRPCRouter } from "~/server/trpc";
import { rootRouter } from "~/server/routers/root";
import { scaleRouter } from "~/server/routers/scale";

export const appRouter = createTRPCRouter({
  root: rootRouter,
  scale: scaleRouter,
});

export type AppRouter = typeof appRouter;
