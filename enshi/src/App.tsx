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

import { Router, Route, createBrowserRouter, createRoutesFromElements, RouterProvider } from "react-router-dom";
import MainPage from "./Pages/MainPage/MainPage";

const router = createBrowserRouter(
    createRoutesFromElements(
        <Route path="/" element={<MainPage />}>
            <Route index element={<Text>Cringer path</Text>}/>
        </Route>
    )
);

export default function App() {
    const { t } = useTranslation();

    return (
        <Theme className="h-fit" accentColor="indigo" grayColor="slate">
            <RouterProvider router={router} />
            <ThemePanel />
        </Theme>
    );
}
