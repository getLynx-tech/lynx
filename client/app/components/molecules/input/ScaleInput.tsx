import React, { useEffect, useState } from "react";
import { type SubmitHandler, useForm } from "react-hook-form";
import { useTRPC } from "~/utils/trpc/react";
import { useMutation } from "@tanstack/react-query";

interface ScaleFormInput {
  pixels: number;
  meters: number;
}

interface ScaleInputProps {
  defaultValues?: ScaleFormInput;
}

export default function ScaleInput({ defaultValues }: ScaleInputProps) {
  const trpc = useTRPC();
  const { register, handleSubmit } = useForm<ScaleFormInput>({
    defaultValues: defaultValues,
  });
  const [scale, setScale] = useState<number | null>(null);

  const postScale = useMutation(trpc.scale.postScale.mutationOptions({}));

  const submit: SubmitHandler<ScaleFormInput> = async (data) => {
    const scale = data.meters / data.pixels;
    setScale(scale);
    postScale.mutate({
      meters: data.meters,
      pixels: data.pixels,
    });
  };

  useEffect(() => {
    if (!defaultValues) return;
    const scale = defaultValues.meters / defaultValues.pixels;
    setScale(scale);
  }, [defaultValues]);

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
              {...register("pixels", { required: true, valueAsNumber: true })}
              className="border-gray-6 w-[250px] rounded-md border-1 p-1 text-sm"
              step="0.01"
              type="number"
            />
          </span>
          <span className="flex flex-col">
            <p className="text-sm">Meters (m)</p>
            <input
              {...register("meters", { required: true, valueAsNumber: true })}
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
