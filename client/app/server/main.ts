import { createTRPCRouter } from "~/server/trpc";
import { rootRouter } from "~/server/routers/root";
import { scaleRouter } from "~/server/routers/scale";
import { anchorRouter } from "~/server/routers/anchor";

export const appRouter = createTRPCRouter({
  root: rootRouter,
  scale: scaleRouter,
  anchorRouter: anchorRouter,
});

export type AppRouter = typeof appRouter;
