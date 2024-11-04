import "./App.css";
import "@radix-ui/themes/styles.css";
import { Theme, ThemePanel } from "@radix-ui/themes";

import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { QueryClientProvider } from "@tanstack/react-query";
import queryClient from "./api/QueryClient/QueryClient";
import { routes } from "./routes/routes";
import { useEffect } from "react";
import "axios";
import { axiosLocalhost } from "./api/axios/axios";

const router = createBrowserRouter(routes);

export default function App() {
    useEffect(() => {
        let f = async () => {
            let c = await axiosLocalhost.post(
                "/login",
                {
                    nickname: "StasikChess",
                    password: "123456",
                }
            );

            console.log(c.headers);
            console.log(document.cookie);
        };

        f();

    }, []);

    return (
        <Theme className="h-fit" accentColor="indigo" grayColor="slate">
            <QueryClientProvider client={queryClient}>
                <RouterProvider router={router} />
                <ThemePanel />
            </QueryClientProvider>
        </Theme>
    );
}
