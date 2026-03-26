import { publicProcedure } from "~/server/trpc";
import { deviceApiFactory } from "~/api/client";

export const deviceRouter = {
  getAllDevices: publicProcedure.query(async () => {
    return await deviceApiFactory.getAllDevices().then((res) => res.data);
  }),
};
