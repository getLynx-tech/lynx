import { Configuration, RootApiFactory } from "~/api/client/generated";

import { axiosInstance } from "~/api/client/axios.server";
import { serverEnv } from "../../../env.server";

const configuration = new Configuration({
  basePath: `http://${serverEnv.SERVER_BASE_URL}`,
});

export const rootApiFactory = RootApiFactory(
  configuration,
  undefined,
  axiosInstance,
);
