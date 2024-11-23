import { Spinner } from "@radix-ui/themes";
import { useQuery } from "@tanstack/react-query";
import { useSetAtom } from "jotai";
import { Outlet } from "react-router-dom";
import { axiosLocalhost } from "../../api/axios/axios";
import { userAtom } from "../../AtomStore/AtomStore";
import NavBar from "../../Components/NavBar/NavBar";

const REFETCH_INTERVAL_IN_MINUTES = 5;
const RETRY_INTERVAL_IN_SECONDS = 1;

const SECONDS_IN_MINUTE = 60;
const MILLS_IN_SECOND = 1000;

export default function MainPage() {
    const setUserData = useSetAtom(userAtom);

    const { isPending } = useQuery({
        queryKey: ["authKey"],
        queryFn: async () => {
            try {
                const response = await axiosLocalhost.get("/auth/check");
                
                setUserData({
                    isAdmin: response.data["is_admin"],
                    username: response.data["username"],
                });
                return true;
            } catch (error) {
                setUserData(undefined);
                return false;
            }
        },
        refetchInterval:
            REFETCH_INTERVAL_IN_MINUTES * SECONDS_IN_MINUTE * MILLS_IN_SECOND,
        refetchOnWindowFocus: true,
        refetchOnReconnect: true,
        gcTime: 10,

        retry: 3,
        retryDelay: (attempt) =>
            attempt * RETRY_INTERVAL_IN_SECONDS * MILLS_IN_SECOND,
    });

    return (
        <>
            {isPending ? (
                <div
                    className="absolute top-1/2 left-1/2 
                                translate-x-[-50%] translate-y-[-50%]"
                >
                    <Spinner size={"3"} />
                </div>
            ) : (
                <>
                    <NavBar />
                    <Outlet />
                </>
            )}
        </>
    );
}
