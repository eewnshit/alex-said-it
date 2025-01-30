import React from "react";

interface Props {
    children: React.ReactNode;
    weight?: "ultra-light" | "thin" | "regular" | "medium" | "semibold" | "bold" | "black" | "heavy";
    size?: number;
    style?: "normal" | "italic",
    display?: "block" | "inline",
    className?: string;
}

const fontWeightMap: { [key: string]: number } = {
    "ultra-light": 100,
    "thin": 200,
    "regular": 400,
    "medium": 600,
    "semibold": 600,
    "bold": 800,
    "black": 900,
    "heavy": 800,
};

export default function Text({ children, className, weight = "regular", size = 12, style = "normal", display = "block" }: Props) {
    const fontWeight = fontWeightMap[weight] || 400;
    const fontSizeStyle = { fontSize: `${size}px` };

    const fontFamily = 'sf-apple';

    return (
        <p
            className={className}
            style={{
                ...fontSizeStyle,
                fontFamily: fontFamily,
                fontWeight: fontWeight,
                fontStyle: style,
                display
            }}
        >
            {children}
        </p>
    );
}
