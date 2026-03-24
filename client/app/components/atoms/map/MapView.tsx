import React, { useEffect, useRef } from "react";
import mapboxgl from "mapbox-gl";
import MapboxDraw from "@mapbox/mapbox-gl-draw";

import "mapbox-gl/dist/mapbox-gl.css";
import "@mapbox/mapbox-gl-draw/dist/mapbox-gl-draw.css";

interface MapViewProps {
  mapBoxAccessToken: string;
}

export default function MapView({ mapBoxAccessToken }: MapViewProps) {
  const mapContainerRef = useRef<HTMLDivElement | null>(null);
  const mapRef = useRef<mapboxgl.Map | null>(null);

  useEffect(() => {
    if (!mapContainerRef.current) return;

    mapboxgl.accessToken = mapBoxAccessToken;

    mapRef.current = new mapboxgl.Map({
      container: mapContainerRef.current,
      style: "mapbox://styles/mapbox/standard",
      center: [11.6532, 48.2489],
      zoom: 18,
    });

    const draw = new MapboxDraw({
      displayControlsDefault: false,
      controls: {
        polygon: true,
        trash: true,
      },
      defaultMode: "draw_polygon",
      styles: [
        {
          id: "gl-draw-polygon-fill",
          type: "fill",
          paint: {
            "fill-color": "#3358D4",
            "fill-opacity": 0.5,
          },
        },
        {
          id: "gl-draw-polygon-stroke",
          type: "line",
          paint: {
            "line-color": "#3358D4",
            "line-width": 3,
          },
        },
        {
          id: "gl-draw-polygon-and-line-vertex",
          type: "circle",
          paint: {
            "circle-radius": 5,
            "circle-color": "#3358D4",
          },
        },
      ],
    });
    mapRef.current.addControl(draw);

    function update() {
      const data = draw.getAll();

      if (data.features.length > 0) {
        const feature = data.features[0];

        if (
          feature.geometry.type === "Polygon" ||
          feature.geometry.type === "MultiPolygon"
        ) {
          const vertices = feature.geometry.coordinates[0];
          console.log("Vertices:", vertices);
        }
      }
    }

    mapRef.current.on("draw.create", update);
    mapRef.current.on("draw.update", update);
    mapRef.current.on("draw.delete", update);
  }, []);

  return <div ref={mapContainerRef} id="map" className="h-full w-full" />;
}
