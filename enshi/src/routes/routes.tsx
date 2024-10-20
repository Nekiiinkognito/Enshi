import { createRoutesFromElements, Route, useRouteError } from "react-router-dom"
import MainPage from "../Pages/MainPage/MainPage"
import {Text} from "@radix-ui/themes";


function ErrorBoundary() {
    let error = useRouteError();
    console.error(error);

    return <div>Dang! This route does not exist... Yet ;)</div>;
}

export const routes = createRoutesFromElements(
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