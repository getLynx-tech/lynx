import {
  AnchorApiFactory,
  Configuration,
  DeviceApiFactory,
  RootApiFactory,
  ScaleApiFactory,
} from "~/api/client/generated";

import { axiosInstance } from "~/api/client/axios.server";
import { serverEnv } from "~/env.server";

const configuration = new Configuration({
  basePath: `http://${serverEnv.SERVER_BASE_URL}`,
});

export const rootApiFactory = RootApiFactory(
  configuration,
  undefined,
  axiosInstance,
);

export const scaleApiFactory = ScaleApiFactory(
  configuration,
  undefined,
  axiosInstance,
);

export const anchorApiFactory = AnchorApiFactory(
  configuration,
  undefined,
  axiosInstance,
);

export const deviceApiFactory = DeviceApiFactory(
  configuration,
  undefined,
  axiosInstance,
);
