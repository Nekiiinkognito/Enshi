import "./App.css";
import "@radix-ui/themes/styles.css";
import {
    Badge,
    Button,
    Callout,
    Container,
    Flex,
    Separator,
    Text,
    Theme,
    ThemePanel,
} from "@radix-ui/themes";

import { useTranslation } from "react-i18next";
import { useRef, useState } from "react";
import parse from "html-react-parser";
import Editor from "./Components/Editor/Editor";

import {
    Router,
    Route,
    createBrowserRouter,
    createRoutesFromElements,
    RouterProvider,
    Routes,
    useRouteError,
} from "react-router-dom";
import MainPage from "./Pages/MainPage/MainPage";
import { QueryClientProvider } from "@tanstack/react-query";
import queryClient from "./api/QueryClient/QueryClient";

function ErrorBoundary() {
    let error = useRouteError();
    console.error(error);

    return <div>Dang! This route does not exist... Yet ;)</div>;
}

const router = createBrowserRouter(
    createRoutesFromElements(
        <>
            <Route
                path="/"
                errorElement={<ErrorBoundary />}
                element={<MainPage />}
            >
                <Route index element={<Text>Cringer path</Text>} />
                <Route
                    path="/a?/c"
                    element={<Text>Cringer path, but this a</Text>}
                ></Route>
            </Route>
        </>
    )
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
