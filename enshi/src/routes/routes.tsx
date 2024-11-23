import { Text } from "@radix-ui/themes";
import {
    createRoutesFromElements,
    Route,
    useRouteError,
} from "react-router-dom";
import AuthPageWrapper from "../Pages/AuthPageWrapper/AuthPageWrapper";
import LoginPage from "../Pages/LoginRegisterPage/LoginPage/LoginPage";
import RegisterPage from "../Pages/LoginRegisterPage/RegisterPage/RegisterPage";
import MainPage from "../Pages/MainPage/MainPage";
import PostCreatorPage from "../Pages/PostCreatorPage/PostCreatorPage";

function ErrorBoundary() {
    let error = useRouteError();
    console.error(error);

    return <div>Dang! This route does not exist... Yet ;)</div>;
}

export const routes = createRoutesFromElements(
    <>
        <Route path="/" errorElement={<ErrorBoundary />} element={<MainPage />}>
            <Route index element={<Text size={"5"}>Cringer path</Text>} />
            <Route
                path="/a?/c"
                element={
                    <Text weight={"regular"}>Cringer path, but this a</Text>
                }
            ></Route>

            <Route
                path="/create"
                element={
                    <AuthPageWrapper>
                        <PostCreatorPage />
                    </AuthPageWrapper>
                }
            ></Route>
        </Route>

        <Route
            path="/login"
            errorElement={<ErrorBoundary />}
            element={<LoginPage />}
        />

        <Route
            path="/register"
            errorElement={<ErrorBoundary />}
            element={<RegisterPage />}
        />
    </>
);
