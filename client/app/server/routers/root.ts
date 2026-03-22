import { publicProcedure } from "~/server/trpc";
import { rootApiFactory } from "~/api/client";

export const rootRouter = {
  getRoot: publicProcedure.query(async ({ ctx }) => {
    return await rootApiFactory.getRoot().then((res) => res.data);
  }),
};
