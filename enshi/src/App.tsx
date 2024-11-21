import { Theme } from "@radix-ui/themes";
import "@radix-ui/themes/styles.css";
import { QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
import "axios";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import queryClient from "./api/QueryClient/QueryClient";
import "./App.css";
import { routes } from "./routes/routes";

const router = createBrowserRouter(routes);

export default function App() {
    return (
        <Theme className="h-fit" accentColor="indigo" grayColor="slate">
            <QueryClientProvider client={queryClient}>
                <RouterProvider router={router} />
                {/* <ThemePanel /> */}
                <ReactQueryDevtools />
            </QueryClientProvider>
        </Theme>
    );
}
