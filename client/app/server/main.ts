import { createTRPCRouter } from "~/server/trpc";
import { rootRouter } from "~/server/routers/root";
import { scaleRouter } from "~/server/routers/scale";
import { anchorRouter } from "~/server/routers/anchor";
import { deviceRouter } from "~/server/routers/device";

export const appRouter = createTRPCRouter({
  root: rootRouter,
  scale: scaleRouter,
  anchorRouter: anchorRouter,
  deviceRouter: deviceRouter,
});

export type AppRouter = typeof appRouter;
