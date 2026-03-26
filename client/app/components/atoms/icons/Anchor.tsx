import * as React from "react";
import type { SVGProps } from "react";
const SvgAnchor = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeLinecap="round"
    strokeLinejoin="round"
    strokeWidth={2}
    className="anchor_svg__lucide anchor_svg__lucide-anchor-icon anchor_svg__lucide-anchor"
    {...props}
  >
    <path d="M12 6v16M19 13l2-1a9 9 0 0 1-18 0l2 1M9 11h6" />
    <circle cx={12} cy={4} r={2} />
  </svg>
);
export default SvgAnchor;
