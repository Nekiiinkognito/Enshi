/** @type {import('tailwindcss').Config} */
export default {
    content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
    theme: {
        extend: {
            colors: {
                "primary-color": "var(--primary-color)",
                "secondary-color": "var(--secondary-color)",
            },
            fontFamily: {
                'times': "Times New Roman"
            },
            animation: {
                appear: "appear 0.25s",
                widthOut: "widthOut cubic-bezier(0.4, 0, 0.6, 1) 0.4s",
                slideFromRight: "slideFromRight cubic-bezier(0.4, 0, 0.6, 1) 0.2s",
                fadeOut: "fadeOut 0.2s ease-in",
            },
            keyframes: {
                fadeOut: {
                    from: {
                        opacity: "1",
                    },
                    to: {
                        opacity: "0",
                    }
                },
                slideFromRight: {
                    "0%": {
                        transform: "translateX(110%)"
                    },
                    "100%": {
                        transform: "translateX(0%)"
                    }
                },
                appear: {
                    "0%": { opacity: "0" },
                    "100%": { opacity: "1" },
                },
                widthOut: {
                    "0%": {
                        width: "0%",
                        left: "50%",
                    },
                    "100%": {
                        width: "100%",
                        left: "0%",
                    },
                },
            },
        },
    },
    plugins: [],
};
