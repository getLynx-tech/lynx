import { publicProcedure } from "~/server/trpc";
import { anchorApiFactory } from "~/api/client";
import z from "zod";

export const anchorRouter = {
  upsertAnchors: publicProcedure
    .input(
      z.array(
        z.object({
          id: z.string(),
          x: z.number(),
          y: z.number(),
        }),
      ),
    )
    .mutation(async ({ input }) => {
      return await anchorApiFactory
        .upsertAnchors({
          anchors: input.map((anchor) => {
            return {
              id: anchor.id,
              x: anchor.x,
              y: anchor.y,
            };
          }),
        })
        .then((res) => res.data);
    }),
  getAllAnchors: publicProcedure.query(async () => {
    return await anchorApiFactory.getAllAnchors().then((res) => res.data);
  }),
};
