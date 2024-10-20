import "./App.css";
import "@radix-ui/themes/styles.css";
import {
    Theme,
    ThemePanel,
} from "@radix-ui/themes";

import {
    createBrowserRouter,
    RouterProvider,
} from "react-router-dom";
import { QueryClientProvider } from "@tanstack/react-query";
import queryClient from "./api/QueryClient/QueryClient";
import { routes } from "./routes/routes";

const router = createBrowserRouter(
    routes
);

export default function App() {
    return (
        <Theme className="h-fit" accentColor="indigo" grayColor="slate">
            <QueryClientProvider client={queryClient}>
                <RouterProvider router={router} />
                <ThemePanel />
            </QueryClientProvider>
        </Theme>
    );
}
