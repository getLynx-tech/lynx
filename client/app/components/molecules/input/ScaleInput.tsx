import React, { useState } from "react";
import { type SubmitHandler, useForm } from "react-hook-form";

interface ScaleInputProps {
  pixels: number;
  meters: number;
}

export default function ScaleInput() {
  const { register, handleSubmit } = useForm<ScaleInputProps>();
  const [scale, setScale] = useState<number | null>(null);

  const submit: SubmitHandler<ScaleInputProps> = async (data) => {
    setScale(data.meters / data.pixels);
  };

  return (
    <form
      onSubmit={handleSubmit(submit)}
      className="border-gray-6 flex flex-col gap-2 rounded-md border-1 bg-white p-2 shadow-md"
    >
      <div className="flex flex-col gap-2">
        <div>
          <span className="flex w-full justify-between">
            <h1 className="text-sm font-semibold text-gray-900">
              Scale Calibration
            </h1>
            <p className="bg-gray-4 cursor-default rounded-md p-0.5 px-1 text-xs">
              Scale: {scale ? Math.round(scale * 100) / 100 : "N/A"}
            </p>
          </span>
          <p className="text-xs text-gray-500">
            Define the real-world distance.
          </p>
        </div>
        <span className="flex flex-col gap-2">
          <span className="flex flex-col">
            <p className="text-sm">Pixels (px)</p>
            <input
              {...register("pixels", { required: true })}
              className="border-gray-6 w-[250px] rounded-md border-1 p-1 text-sm"
              step="0.01"
              type="number"
            />
          </span>
          <span className="flex flex-col">
            <p className="text-sm">Meters (m)</p>
            <input
              {...register("meters", { required: true })}
              type="number"
              step="0.01"
              className="border-gray-6 w-[250px] rounded-md border-1 p-1 text-sm"
            />
          </span>
        </span>
      </div>
      <button className="cursor-pointer rounded-md bg-black px-2 py-1 text-sm text-white">
        Submit
      </button>
    </form>
  );
}
