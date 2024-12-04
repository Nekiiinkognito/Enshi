import { Text } from "@radix-ui/themes";
import {
    createRoutesFromElements,
    Route,
    useRouteError,
} from "react-router-dom";
import ArticleViewer from "../Components/ArticleViewer/ArticleViewer";
import AuthPageWrapper from "../Pages/AuthPageWrapper/AuthPageWrapper";
import LoginPage from "../Pages/LoginRegisterPage/LoginPage/LoginPage";
import PostRedactor from "../Pages/LoginRegisterPage/PostRedactor/PostRedactor";
import RegisterPage from "../Pages/LoginRegisterPage/RegisterPage/RegisterPage";
import MainPage from "../Pages/MainPage/MainPage";
import PostCreatorPage from "../Pages/PostCreatorPage/PostCreatorPage";
import RandomPostsPage from "../Pages/RandomPostsPage/RandomPostsPage";

function ErrorBoundary() {
    let error = useRouteError();
    console.error(error);

    return <div>Dang! This route does not exist... Yet ;)</div>;
}

export const routes = createRoutesFromElements(
    <>
        <Route path="/" errorElement={<ErrorBoundary />} element={<MainPage />}>
            <Route index element={<RandomPostsPage />} />

            <Route
                path="/a?/c"
                element={
                    <Text weight={"regular"}>This page is yet to be created</Text>
                }
            />

            <Route
                path="/create"
                element={
                    <AuthPageWrapper>
                        <PostCreatorPage />
                    </AuthPageWrapper>
                }
            />

            <Route path="/posts/:postId" element={<ArticleViewer />} />
            <Route path="/posts/change/:postId" element={<PostRedactor />} />
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
