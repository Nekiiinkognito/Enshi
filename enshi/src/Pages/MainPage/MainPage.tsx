import React from "react";
import { Outlet } from "react-router-dom";
import NavBar from "../../Components/NavBar/NavBar";
import { axiosLocalhost } from "../../api/axios/axios";
import { Container } from "@radix-ui/themes";

export default function MainPage() {
    return (
        <>
            <NavBar />
            <Outlet />
            <button
                onClick={async () => {
                    let d = await axiosLocalhost.get("getCookie");
                    console.log(d.data);
                }}
            >
                Click for cookie test
            </button>
        </>
    );
}
