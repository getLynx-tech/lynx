import { publicProcedure } from "~/server/trpc";
import { scaleApiFactory } from "~/api/client";
import { z } from "zod";

export const scaleRouter = {
  getScale: publicProcedure.query(async () => {
    return await scaleApiFactory.getScale().then((res) => res.data);
  }),
  postScale: publicProcedure
    .input(
      z.object({
        meters: z.number(),
        pixels: z.number(),
      }),
    )
    .mutation(async ({ input }) => {
      return await scaleApiFactory
        .createScale({
          meters: input.meters,
          pixels: input.pixels,
        })
        .then((res) => res.data);
    }),
};
